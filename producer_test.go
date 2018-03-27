package hydra

import "testing"

func defaultProducer() *Producer {
	config := DefaultConfig()

	return NewProducer(config)
}

func TestNewProducer(t *testing.T) {
	producer := defaultProducer()
	if len(producer.topics) != 0 {
		t.Error("Producer should have 0 topics by default")
	}
}
