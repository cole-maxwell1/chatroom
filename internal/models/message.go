package models

import "time"

type ChatMessage struct {
	Content   string
	Username  string
	Timestamp time.Time
}
