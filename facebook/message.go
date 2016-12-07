package facebook

// Message represents a single message
type Message struct {
	Sender   string `json:"sender"`
	Metadata string `json:"meta"`
	Body     string `json:"body"`
}

// NewMessage creates a new message
func NewMessage(sender string, meta string, body string) *Message {
	return &Message{Sender: sender, Metadata: meta, Body: body}
}
