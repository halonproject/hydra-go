package hydra

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	ipfs "github.com/ipfs/go-ipfs-api"
)

func ipfsIsRunning() bool {
	cmd := exec.Command("lsof", "-n", "-i", ":5001")
	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}

func defaultConsumer() (*Consumer, error) {
	config := DefaultConfig()

	return NewConsumer(newMockIPFSClient(), config)
}

func productionConsumer() (*Consumer, error) {
	if !ipfsIsRunning() {
		return nil, fmt.Errorf("local IPFS node is not available")
	}

	config := DefaultConfig()
	shell := ipfs.NewShell(fmt.Sprintf("%s:%s", config.IPFSAddr, config.IPFSPort))
	return NewConsumer(shell, config)
}

func TestNewConsumer(t *testing.T) {
	_, err := defaultConsumer()
	if err != nil {
		t.Error("Error making default consumer:", err.Error())
	}
}

func TestConsumerTopics(t *testing.T) {
	consumer, err := defaultConsumer()
	if err != nil {
		t.Error("Error making default consumer:", err.Error())
	}

	if len(consumer.Topics()) != 0 {
		t.Error("Consumer should not have any subscribed topics")
	}

	consumer.SubscribeTopics([]string{"topic1", "topic2"})

	if numTopics := len(consumer.Topics()); numTopics != 2 {
		t.Error("Consumer should have 2 subscribed topics, but got ", numTopics)
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

// All integration testing functions

func TestConsumeStart(t *testing.T) {
	consumer, err := productionConsumer()
	if err != nil {
		t.Log("Error making production consumer:", err.Error())
		return
	}

	consumer.SubscribeTopics([]string{"foo", "bar"})

	consumer.Start()

	time.Sleep(time.Second * 5)

	consumer.Stop()
}

func TestUnsubscribe(t *testing.T) {
	consumer, err := productionConsumer()
	if err != nil {
		t.Log("Error making production consumer:", err.Error())
		return
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

func TestReadMessage(t *testing.T) {
	consumer, err := productionConsumer()
	if err != nil {
		t.Log("Error making production consumer:", err.Error())
		return
	}

	producer, err := productionProducer()
	if err != nil {
		t.Log("Error making production producer:", err.Error())
		return
	}

	err = consumer.Subscribe("foo")
	if err != nil {
		t.Error("Error subscribing to topic:", err.Error())
	}

	consumer.Start()
	go func() {
		time.Sleep(time.Second * 5)
		producer.AddTopic("foo")
		msg := NewMessage([]byte("test"), []byte("test_message"), nil)

		producer.ProduceAll(msg)
	}()

	_, err = consumer.ReadMessage()
	if err != nil {
		t.Error("Error reading message for topic 'foo'", err.Error())
	}
}
