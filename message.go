package hydra

import (
	"fmt"
	"time"

	message "github.com/halonproject/hydra-proto-go"
)

// Message is a generic message that is sent through a topic via IPFS pubsub.
type Message struct {
	Value     []byte            `json:"value"`
	Key       []byte            `json:"key"`
	Timestamp int64             `json:"timestamp"`
	Headers   []*message.Header `json:"headers"`
}

// NewMessage creates a new message with key, value and any Headers needed
func NewMessage(key, value []byte, headers []*message.Header) *message.Message {
	return &message.Message{
		Key:       key,
		Value:     value,
		Headers:   headers,
		Timestamp: time.Now().Unix(),
	}
}

// String returns a string representation of a message to conform to Event interface
func (msg *Message) String() string {
	return fmt.Sprintf("%d: %s: %s", msg.Timestamp, msg.Key, msg.Value)
}
