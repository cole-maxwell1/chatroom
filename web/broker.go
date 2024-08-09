// Credit to the Gorilla Websocket Chat Example:
// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// See: https://github.com/gorilla/websocket/tree/main/examples/chat

/*
@cole-maxwell1
My modified version of the Gorilla Websocket Chat Example
just formats and censors messages before broadcasting them.
Select functionality moved to broker functions.
*/

package web

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/TwiN/go-away"
	"github.com/cole-maxwell1/chatroom/internal/models"
	"github.com/cole-maxwell1/chatroom/internal/pkg"
	"github.com/cole-maxwell1/chatroom/web/templates"
)

const (
	MAX_MESSAGES = 100
)

type InboundMessage struct {
	UnsanitizedMessage []byte
	Client             *Client
}

type WebSocketBroker struct {
	clients      map[*Client]bool
	inbound      chan InboundMessage
	register     chan *Client
	unregister   chan *Client
	chatMessages *pkg.RingBuffer[models.ChatMessage]
}

func NewHub() *WebSocketBroker {
	return &WebSocketBroker{
		clients:      make(map[*Client]bool),
		inbound:      make(chan InboundMessage),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		chatMessages: pkg.NewRingBuffer[models.ChatMessage](MAX_MESSAGES),
	}
}

func (b *WebSocketBroker) Run() {
	for {
		select {
		case client := <-b.register:
			b.registerClient(client)
		case client := <-b.unregister:
			b.unregisterClient(client)
		case message := <-b.inbound:
			b.handleIncomingMsg(message)
		}
	}
}

type NewMessage struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

func (b *WebSocketBroker) handleIncomingMsg(msg InboundMessage) {
	var incomingMsg NewMessage
	err := json.Unmarshal(msg.UnsanitizedMessage, &incomingMsg)
	if err != nil {
		msg.Client.send <- []byte("Invalid message format. Please ensure your message is properly formatted JSON.")
		return
	}

	sanitizedMsg := goaway.Censor(incomingMsg.Message)
	sanitizedUsername := goaway.Censor(incomingMsg.Username)

	formattedMessage := models.ChatMessage{
		Content:   sanitizedMsg,
		Username:  sanitizedUsername,
		Timestamp: time.Now(),
	}
	b.chatMessages.Add(formattedMessage)

	var templBytes bytes.Buffer
	templates.ChatMessageSwap(formattedMessage).Render(context.TODO(), &templBytes)

	b.broadcastMessage(templBytes.Bytes())
}

func (b *WebSocketBroker) registerClient(client *Client) {
	b.clients[client] = true
	b.updateChatterTotal()
}

func (b *WebSocketBroker) unregisterClient(client *Client) {
	if _, ok := b.clients[client]; ok {
		delete(b.clients, client)
		close(client.send)
		b.updateChatterTotal()
	}
}

func (b *WebSocketBroker) broadcastMessage(message []byte) {
	for client := range b.clients {
		select {
		case client.send <- message:
		default:
			b.unregisterClient(client)
		}
	}
}

func (b *WebSocketBroker) updateChatterTotal() {
	totalChatters := len(b.clients)

	// Render the updated chat display
	var templBytes bytes.Buffer
	templates.TotalChattersSwap(totalChatters).Render(context.TODO(), &templBytes)

	// Broadcast the updated chat display
	b.broadcastMessage(templBytes.Bytes())
}
