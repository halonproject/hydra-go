package hydra

import (
	"testing"
)

func TestMessageString(t *testing.T) {
	header := NewHeader("hello", []byte("world"))
	message := NewMessage([]byte("foo"), []byte("bar"), []Header{header})

	if str := message.String(); str == "" {
		t.Error("Expected non-emtpy string for message string, but got", str)
	}
}

func TestMessageEncode(t *testing.T) {
	header := NewHeader("hello", []byte("world"))
	message := NewMessage([]byte("foo"), []byte("bar"), []Header{header})

	msgBytes, err := message.Encode()
	if err != nil {
		t.Error("Error encoding message")
	}

	t.Log(msgBytes)
}

func TestMessageDecodeJSONMap(t *testing.T) {
	header := NewHeader("Content-Type", []byte("application/json"))
	message := NewMessage([]byte("json_message"), []byte(`{"key":"value"}`), []Header{header})

	obj, err := message.decodeJSON()
	if err != nil {
		t.Error("Error decoding message:", err.Error())
	}

	switch obj.(type) {
	case map[string]string:
		return
	case nil:
		t.Error("Should have recieved non-nil interface")
	}
}

func TestMessageDecodeJSONArray(t *testing.T) {
	header := NewHeader("Content-Type", []byte("application/json"))
	message := NewMessage([]byte("json_message"), []byte(`[{"key":"value"}]`), []Header{header})

	obj, err := message.decodeJSON()
	if err != nil {
		t.Error("Error decoding message:", err.Error())
	}

	switch obj.(type) {
	case []map[string]string:
		return
	case nil:
		t.Error("Should have recieved non-nil interface")
	}
}

func TestMessageDecodeMap(t *testing.T) {
	header := NewHeader("Content-Type", []byte("application/json"))
	message := NewMessage([]byte("json_message"), []byte(`{"key":"value"}`), []Header{header})

	obj, err := message.Decode()
	if err != nil {
		t.Error("Error decoding message:", err.Error())
	}

	switch obj.(type) {
	case []map[string]string:
		return
	case nil:
		t.Error("Should have recieved non-nil interface")
	}
}

func TestMessageDecodeArray(t *testing.T) {
	header := NewHeader("Content-Type", []byte("application/json"))
	message := NewMessage([]byte("json_message"), []byte(`[{"key":"value"}]`), []Header{header})

	obj, err := message.Decode()
	if err != nil {
		t.Error("Error decoding message:", err.Error())
	}

	switch obj.(type) {
	case []map[string]string:
		return
	case nil:
		t.Error("Should have recieved non-nil interface")
	}
}
