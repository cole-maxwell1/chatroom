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
	"time"

	"github.com/a-h/templ"
	"github.com/cole-maxwell1/chatroom/internal/models"
	"github.com/cole-maxwell1/chatroom/internal/pkg"
	"github.com/cole-maxwell1/chatroom/web/templates"
)

// broker maintains the set of active clients and broadcasts messages to the
// clients.
type WebSocketBroker struct {
	// Registered clients.
	clients map[*Client]bool

	inbound chan []byte

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
		inbound:          make(chan []byte),
		broadcast:        make(chan []byte),
		register:         make(chan *Client),
		unregister:       make(chan *Client),
		clients:          make(map[*Client]bool),
		connectionChange: make(chan struct{}, 1), // Must be Buffered channel so will not block if there isn't an immediate receiver
	}
}

func (h *WebSocketBroker) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.connectionChange <- struct{}{}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.connectionChange <- struct{}{}
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.inbound:
			go func(msg []byte) {
				sanitizedMessage := sanitizeMessage(msg)
				h.broadcast <- sanitizedMessage
			}(message)

		case <-h.connectionChange:
			go func(tot int) {
				h.broadcast <- renderTotalChatters(tot)

			}(len(h.clients))
		}
	}
}

func renderTotalChatters(numChatters int) []byte {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	templates.TotalChatters(numChatters).Render(context.Background(), buf)
	return buf.Bytes()
}

const MAX_MESSAGES = 100

var chatMessages = pkg.NewRingBuffer[models.ChatMessage](MAX_MESSAGES)

func sanitizeMessage(msg []byte) []byte {

	// Remove leading and trailing whitespace from the message
	msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))

	newMessage := models.ChatMessage{
		Content:           string(msg),
		Username:          string(' '), // Replace the empty rune literal with a space rune
		FormattedDateTime: pkg.FormatDate(time.Now()),
	}

	chatMessages.Add(newMessage)

	var templBytes bytes.Buffer
	templates.ChatMessage(newMessage).Render(context.Background(), &templBytes)

	return templBytes.Bytes()
}
