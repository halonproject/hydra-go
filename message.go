package hydra

import (
	"fmt"
	"time"
)

type Message struct {
	Value     []byte
	Key       []byte
	Timestamp time.Time
	Headers   []Header
}

func NewMessage(key, value []byte, headers []Header) *Message {
	return &Message{}
}

func (msg *Message) String() string {
	return fmt.Sprintf("%s: [%s=%s]", msg.Timestamp, msg.Key, msg.Value)
}
