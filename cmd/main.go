package main

import (
	"fmt"

	"github.com/cole-maxwell1/chatroom/internal/pkg"
)

// Message represents a chat message
type Message struct {
	Content           string
	Username          string
	FormattedDateTime string
}

func main() {
	const TOTAL_STORED_MESSAGES = 500
	messageBuffer := pkg.NewRingBuffer[Message](TOTAL_STORED_MESSAGES)

	for i := 1; i <= 501; i++ {
		messageBuffer.Add(Message{Content: fmt.Sprintf("Message %d", i)})
	}

	messages := messageBuffer.Get()
	for _, msg := range messages {
		fmt.Println(msg.Content)
	}
}
