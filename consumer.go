package hydra

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type Consumer struct {
	events       chan Event
	topics       []string
	conn         *ipfs.Shell
	subscription *ipfs.PubSubSubscription
}

func NewConsumer(config *Config) (*Consumer, error) {
	consumer := &Consumer{}
	dsn := fmt.Sprintf("https://%s:%s", config.IPFSAddr, config.IPFSPort)

	shell := ipfs.NewShell(dsn)
	consumer.conn = shell
	consumer.subscription = &ipfs.PubSubSubscription{}
	consumer.topics = config.Topics
	consumer.events = make(chan Event)

	return consumer, nil
}

func (c *Consumer) Subscribe(topic string) error {
	return c.SubscribeTopics([]string{topic})
}

func (c *Consumer) SubscribeTopics(topics []string) error {
	c.topics = append(c.topics, topics...)
	return nil
}

func (c *Consumer) Unsubscribe(topic string) error {
	return c.UnsubscribeTopics([]string{topic})
}

func (c *Consumer) UnsubscribeTopics(topics []string) error {
	newTopics := make([]string, 0)
	for _, topic := range topics {
		for _, ctopic := range c.topics {
			if topic != ctopic {
				newTopics = append(newTopics, ctopic)
			}
		}
	}

	c.topics = newTopics
	return nil
}

func (c *Consumer) ReadMessage(timeout time.Duration) (*Message, error) {
	var timeoutMs int

	if timeout > 0 {
		timeoutMs = (int)(timeout.Seconds() * 1000.0)
	} else {
		timeoutMs = (int)(timeout)
	}
	for {
		event := c.Poll(timeoutMs)

		switch e := event.(type) {
		case *Message:
			return e, nil
		case *Error:
			return nil, errors.New(e.reason)
		default:
			// ignore everything else
		}

		if timeoutMs == 0 && event == nil {
			return nil, newError(HYDRA_RESPONSE_TIMEOUT_ERROR)
		}
	}
}

func (c *Consumer) Poll(timeoutMs int) Event {
	return <-c.events
}

func (c *Consumer) consumeAllTopics() {
	for {
		record, err := c.subscription.Next()
		if err == nil && sliceContainsSliceElement(c.topics, record.TopicIDs()) {
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(record.SeqNo()))
			c.events <- &Message{Key: b, Value: record.Data()}
		}
	}
}

func sliceContainsSliceElement(haystack, needles []string) bool {
	elementMap := make(map[string]bool)
	for _, e := range haystack {
		elementMap[e] = true
	}

	for _, str := range needles {
		if elementMap[str] {
			return true
		}
	}
	return false
}
