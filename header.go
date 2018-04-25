package hydra

import (
	message "github.com/halonproject/hydra-proto-go"
)

// Header represents a key-value pair that stores meta data about a message.
type Header struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

// NewHeader creates a new Header with key and value
func NewHeader(key string, value []byte) *message.Header {
	return &message.Header{Key: key, Value: value}
}
