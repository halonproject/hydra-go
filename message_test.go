package hydra

import "testing"

func TestMessageString(t *testing.T) {
	header := NewHeader("hello", []byte("world"))
	message := NewMessage([]byte("foo"), []byte("bar"), []Header{header})

	if str := message.String(); str == "" {
		t.Error("Expected non-emtpy string for message string, but got", str)
	}
}
