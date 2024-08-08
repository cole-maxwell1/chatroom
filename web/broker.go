// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// See: https://github.com/gorilla/websocket/tree/main/examples/chat

/*
Cole Maxwell 2024

Modifications:
- Added a new channel to the broker struct called connectionChange.
This channel is used to notify the broker when a client connects or disconnects.
- Add and additional channel for inbound messages from the clients.
This channel is used to intercept messages to sanitize and validate them
before they are broadcasted to all clients.
*/

package web

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"time"

	goaway "github.com/TwiN/go-away"
	"github.com/a-h/templ"
	"github.com/cole-maxwell1/chatroom/internal/models"
	"github.com/cole-maxwell1/chatroom/internal/pkg"
	"github.com/cole-maxwell1/chatroom/web/templates"
)

type InboundMessage struct {
	UnsanitizedMessage []byte
	Client             *Client
}

// broker maintains the set of active clients and broadcasts messages to the
// clients.
type WebSocketBroker struct {
	// Registered clients.
	clients map[*Client]bool

	inbound chan InboundMessage

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	connectionChange chan struct{}
}

func NewHub() *WebSocketBroker {
	return &WebSocketBroker{
		inbound:          make(chan InboundMessage),
		broadcast:        make(chan []byte),
		register:         make(chan *Client),
		unregister:       make(chan *Client),
		clients:          make(map[*Client]bool),
		connectionChange: make(chan struct{}, 1), // Must be Buffered channel so will not block if there isn't an immediate receiver
	}
}

func (b *WebSocketBroker) Run() {
	for {
		select {
		case client := <-b.register:
			b.clients[client] = true
			b.connectionChange <- struct{}{}
		case client := <-b.unregister:
			if _, ok := b.clients[client]; ok {
				delete(b.clients, client)
				close(client.send)
				b.connectionChange <- struct{}{}
			}
		case message := <-b.broadcast:
			for client := range b.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(b.clients, client)
				}
			}
		case message := <-b.inbound:
			go b.sanitizeMessage(message)
		case <-b.connectionChange:
			go func(tot int) {
				b.broadcast <- renderTotalChatters(tot)
			}(len(b.clients))
		}
	}
}

func renderTotalChatters(numChatters int) []byte {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	templates.TotalChattersSwap(numChatters).Render(context.Background(), buf)
	return buf.Bytes()
}

const MAX_MESSAGES = 100

var chatMessages = pkg.NewRingBuffer[models.ChatMessage](MAX_MESSAGES)

type NewMessage struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

func (b *WebSocketBroker) sanitizeMessage(msg InboundMessage) {

	var incomingMsg NewMessage
	// Parse msg to json object
	err := json.Unmarshal(msg.UnsanitizedMessage, &incomingMsg)
	if err != nil {
		msg.Client.send <- []byte("Invalid message format")
	} else {
		// Sanitize the message
		sanitizedMsg := goaway.Censor(incomingMsg.Message)
		//remove space from username
		sanitizedUsername := strings.Replace(incomingMsg.Username, " ", "", -1)
		sanitizedMsg = goaway.Censor(sanitizedMsg)

		formattedMessage := models.ChatMessage{
			Content:   sanitizedMsg,
			Username:  sanitizedUsername, // Replace the empty rune literal with a space rune
			Timestamp: time.Now(),
		}

		chatMessages.Add(formattedMessage)

		var templBytes bytes.Buffer
		templates.ChatMessageSwap(formattedMessage).Render(context.Background(), &templBytes)

		b.broadcast <- templBytes.Bytes()
	}

}
