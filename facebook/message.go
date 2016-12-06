package facebook

import "time"

// Message represents a single message
type Message struct {
	Sender    string
	Recipient string
	Metadata  time.Time
	Body      string
}

// NewMessage creates a new message
func NewMessage(sender string, recipient string, meta time.Time, body string) *Message {
	return &Message{Sender: sender, Recipient: recipient, Metadata: meta, Body: body}
}
