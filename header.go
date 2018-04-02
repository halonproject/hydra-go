package hydra

import "fmt"

// Header represents a key-value pair that stores meta data about a message.
type Header struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

// NewHeader creates a new Header with key and value
func NewHeader(key string, value []byte) Header {
	return Header{Key: key, Value: value}
}

// String returns a string representation of a header
func (h Header) String() string {
	return fmt.Sprintf("%s=%s", h.Key, h.Value)
}
