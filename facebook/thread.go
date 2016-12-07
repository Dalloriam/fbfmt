package facebook

import (
	"errors"
)

// Thread represents a thread
type Thread struct {
	Participants []string
	Messages     []Message
}

func (t *Thread) canMerge(other *Thread) bool {

	if t.Participants == nil && other.Participants == nil {
		return true
	}

	if t.Participants == nil || other.Participants == nil {
		return false
	}

	if len(t.Participants) != len(other.Participants) {
		return false
	}

	for i := range t.Participants {
		if t.Participants[i] != other.Participants[i] {
			return false
		}
	}

	return true
}

// Merge merges two thread objects
func (t *Thread) Merge(other *Thread) error {
	if t.canMerge(other) {
		t.Messages = append(t.Messages, other.Messages...)
	} else {
		return errors.New("Threads is not mergeable. (Not the same participants)")
	}

	return nil
}
