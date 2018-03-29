package hydra

import "testing"

func defaultProducer() *Producer {
	config := DefaultConfig()

	return NewProducer(config)
}

func TestSliceContainsString(t *testing.T) {
	needle := "hello"

	haystack := []string{"hello", "world"}

	if !sliceContainsString(haystack, needle) {
		t.Error("Slice should contain string \"hello\"")
	}

	needle = "nope"

	if sliceContainsString(haystack, needle) {
		t.Error("Slice should not contain \"nope\"")
	}
}

func TestNewProducer(t *testing.T) {
	producer := defaultProducer()
	if len(producer.topics) != 0 {
		t.Error("Producer should have 0 topics by default")
	}
}

func TestAddTopics(t *testing.T) {
	producer := defaultProducer()

	producer.AddTopic("foo")

	if len(producer.Topics()) != 1 {
		t.Error("Producer should have one topic")
	}

	producer.AddTopics([]string{"foo", "bar", "baz"})

	if len(producer.Topics()) != 3 {
		t.Error("Producer should have 3 topics")
	}
}

func TestRemoveTopics(t *testing.T) {
	producer := defaultProducer()

	producer.AddTopic("foo")

	if len(producer.Topics()) != 1 {
		t.Error("Producer should have one topic")
	}

	producer.AddTopics([]string{"foo", "bar", "baz"})

	if topics := len(producer.Topics()); topics != 3 {
		t.Error("Producer should have 3 topics but has", topics)
	}

	producer.RemoveTopic("foo")

	if topics := len(producer.Topics()); topics != 2 {
		t.Error("Producer should have 2 topics but has", topics)
	}

	producer.RemoveTopics([]string{"abc", "xyz"})

	if topics := len(producer.Topics()); topics != 2 {
		t.Error("Producer should have 2 topics but has", topics)
	}

	producer.RemoveTopics([]string{"bar", "baz"})

	if topics := len(producer.Topics()); topics != 0 {
		t.Error("Producer should have 0 topics but has", topics)
	}
}
