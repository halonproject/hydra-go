package hydra

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"time"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type Consumer struct {
	events        chan Event
	topics        []string
	conn          *ipfs.Shell
	subscriptions map[string]*ipfs.PubSubSubscription
}

func NewConsumer(config *Config) (*Consumer, error) {
	consumer := &Consumer{}
	dsn := fmt.Sprintf("https://%s:%s", config.IPFSAddr, config.IPFSPort)

	shell := ipfs.NewShell(dsn)
	consumer.conn = shell
	consumer.subscriptions = make(map[string]*ipfs.PubSubSubscription)
	consumer.topics = config.Topics
	consumer.events = make(chan Event)

	return consumer, nil
}

func (c *Consumer) Topics() []string {
	return c.topics
}

func (c *Consumer) Subscribe(topic string) error {
	return c.SubscribeTopics([]string{topic})
}

func (c *Consumer) SubscribeTopics(topics []string) error {
	for _, topic := range topics {
		c.topics = append(c.topics, topic)
		subscription, err := c.conn.PubSubSubscribe(topic)
		if err != nil {
			return err
		}
		c.subscriptions[topic] = subscription
	}

	return nil
}

func (c *Consumer) Unsubscribe(topic string) error {
	return c.UnsubscribeTopics([]string{topic})
}

func (c *Consumer) UnsubscribeTopics(topics []string) error {
	var err error
	for _, topic := range topics {
		for i, ctopic := range c.topics {
			if topic == ctopic {
				c.topics = append(c.topics[:i], c.topics[i+1:]...)
			} else {
				if c.subscriptions[topic] != nil {
					err = c.subscriptions[topic].Cancel()
					delete(c.subscriptions, topic)
				}
			}
		}
	}

	return err
}

func (c *Consumer) ReadMessage(timeout time.Duration) (*Message, error) {
	for {
		event := c.Poll()

		switch e := event.(type) {
		case *Message:
			return e, nil
		case *Error:
			return nil, errors.New(e.reason)
		default:
			// ignore everything else
		}
	}
}

func (c *Consumer) Poll() Event {
	return <-c.events
}

func (c *Consumer) consumeAllTopics() {
	wg := &sync.WaitGroup{}
	for _, subscription := range c.subscriptions {
		go func(subscription *ipfs.PubSubSubscription, wg *sync.WaitGroup) {
			defer wg.Done()
			wg.Add(1)
			c.consumeTopic(subscription)
		}(subscription, wg)
	}
	wg.Wait()
}

func (c *Consumer) consumeTopic(subscription *ipfs.PubSubSubscription) {
	for {
		record, err := subscription.Next()
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
