package facebook

// Thread represents a thread
type Thread struct {
	Participants []string   `json:"participants"`
	Messages     []*Message `json:"messages"`
}
