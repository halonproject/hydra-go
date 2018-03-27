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
