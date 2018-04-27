package hydra

import (
	"time"

	message "github.com/halonproject/hydra-proto-go"
)

// NewMessage creates a new message with key, value and any Headers needed
func NewMessage(key, value []byte, headers []*message.Header) *message.Message {
	return &message.Message{
		Key:       key,
		Value:     value,
		Headers:   headers,
		Timestamp: time.Now().Unix(),
	}
}
