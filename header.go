package hydra

import (
	message "github.com/halonproject/hydra-proto-go"
)

// NewHeader creates a new Header with key and value
func NewHeader(key string, value []byte) *message.Header {
	return &message.Header{Key: key, Value: value}
}
