package hydra

import "testing"

func defaultConsumer() (*Consumer, error) {
	config := DefaultConfig()

	return NewConsumer(config)
}

func TestNewConsumer(t *testing.T) {
	_, err := defaultConsumer()
	if err != nil {
		t.Error("Error making default consumer:", err.Error())
	}
}

func TestSubscribe(t *testing.T) {
	consumer, err := defaultConsumer()
	if err != nil {
		t.Error("Error making default consumer:", err.Error())
	}

	err = consumer.Subscribe("foo")
	if err != nil {
		t.Error("Error subscribing to topic:", err.Error())
	}
}

func TestUnsubscribe(t *testing.T) {
	consumer, err := defaultConsumer()
	if err != nil {
		t.Error("Error making default consumer:", err.Error())
	}

	err = consumer.Subscribe("foo")
	if err != nil {
		t.Error("Error subscribing to topic:", err.Error())
	}

	err = consumer.Unsubscribe("foo")
	if err != nil {
		t.Error("Error unsubscribing from topic", err.Error())
	}
}

func TestSliceContainsSliceElement(t *testing.T) {
	array1 := []string{"one", "two", "three"}
	array2 := []string{"aaa", "bbb", "three"}

	contains := sliceContainsSliceElement(array1, array2)
	if !contains {
		t.Error("haystack should contain needles")
	}

	array2 = []string{"aaa", "bbb", "ccc"}

	contains = sliceContainsSliceElement(array1, array2)
	if contains {
		t.Error("haystack should NOT contain needles")
	}
}
