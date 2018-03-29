package hydra

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Value     []byte   `json:"value"`
	Key       []byte   `json:"key"`
	Timestamp int64    `json:"timestamp"`
	Headers   []Header `json:"headers"`
}

func NewMessage(key, value []byte, headers []Header) *Message {
	return &Message{
		Key:       key,
		Value:     value,
		Headers:   headers,
		Timestamp: time.Now().Unix(),
	}
}

func (msg *Message) String() string {
	return fmt.Sprintf("%d: [%s=%s]", msg.Timestamp, msg.Key, msg.Value)
}

func (msg *Message) Encode() ([]byte, error) {
	return json.Marshal(msg)
}

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
