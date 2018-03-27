package hydra

import "fmt"

type Header struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
}

func NewHeader(key string, value []byte) Header {
	return Header{Key: key, Value: value}
}

func (h Header) String() string {
	return fmt.Sprintf("%s=%s", h.Key, h.Value)
}
