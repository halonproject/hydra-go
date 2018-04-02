package hydra

import (
	"encoding/json"
	"fmt"
	"time"
)

// Message is a generic message that is sent through a topic via IPFS pubsub.
type Message struct {
	Value     []byte   `json:"value"`
	Key       []byte   `json:"key"`
	Timestamp int64    `json:"timestamp"`
	Headers   []Header `json:"headers"`
}

// NewMessage creates a new message with key, value and any Headers needed
func NewMessage(key, value []byte, headers []Header) *Message {
	return &Message{
		Key:       key,
		Value:     value,
		Headers:   headers,
		Timestamp: time.Now().Unix(),
	}
}

// String returns a string representation of a message to conform to Event interface
func (msg *Message) String() string {
	return fmt.Sprintf("%d: [%s=%s]", msg.Timestamp, msg.Key, msg.Value)
}

// Encode return the json marshalled message and an error if one occurred
func (msg *Message) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

// Decode returns a decoded message as an interface. It will attempt to un-encode
// the message value depending on the headers that were sent on the message.
func (msg *Message) Decode() (interface{}, error) {
	for _, header := range msg.Headers {
		key := header.Key
		switch key {
		case "Content-Type":
			value := string(header.Value)
			switch value {
			case "application/json":
				return msg.decodeJSON()
			}
		case "Authorization":
			// TODO: determine how we will want to pass encrypted messages
		}
	}

	return msg.Value, nil
}

// decodeJSON will attempt to take the raw message value and unmarshall it into
// either a map[string]interface{} or []interface or will return an error if it
// could not covert the raw message value into either type
func (msg *Message) decodeJSON() (interface{}, error) {
	var objMap map[string]interface{}
	err := json.Unmarshal(msg.Value, &objMap)
	if err == nil {
		return objMap, nil
	}

	var objArr []interface{}
	err = json.Unmarshal(msg.Value, &objArr)
	if err == nil {
		return objArr, nil
	}

	return nil, err
}
