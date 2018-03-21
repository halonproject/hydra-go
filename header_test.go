package hydra

import "testing"

func TestHeaderString(t *testing.T) {
	header := NewHeader("foo", []byte("bar"))

	if str := header.String(); str != "foo=bar" {
		t.Error("Expected header string to be \"foo=bar\" but got", str)
	}
}
