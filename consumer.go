package hydra

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	ipfs "github.com/ipfs/go-ipfs-api"
)

// Consumer "consumes" all topic messages that it is subscribed to. It provides
// a high level interface for topic management and pulling messages from IPFS
// pubsub.
type Consumer struct {
	events        chan Event
	topics        []string
	conn          *ipfs.Shell
	subscriptions map[string]*ipfs.PubSubSubscription
}

// NewConsumer creates a new consumer that is connected to a IPFS client via the
// configuration passed and also is subscribed to any topics set in the config.
func NewConsumer(config *Config) (*Consumer, error) {
	consumer := &Consumer{}
	dsn := fmt.Sprintf("http://%s:%s", config.IPFSAddr, config.IPFSPort)

	shell := ipfs.NewShell(dsn)
	consumer.conn = shell
	consumer.subscriptions = make(map[string]*ipfs.PubSubSubscription)
	consumer.topics = config.Topics
	consumer.events = make(chan Event, 10000)

	return consumer, nil
}

// Topics returns the current list of topics that the consumer is subscribed to.
func (c *Consumer) Topics() []string {
	return c.topics
}

// Subscribe add a topic to the consumers list of topics and will allow messages
// to be consumed from those topics.
func (c *Consumer) Subscribe(topic string) error {
	return c.SubscribeTopics([]string{topic})
}

// SubscribeTopics adds a list of topics that the consumer will consume messages
// from.
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

// Unsubscribe removes a topic from the list of topics that the conumser is
// consuming.
func (c *Consumer) Unsubscribe(topic string) error {
	return c.UnsubscribeTopics([]string{topic})
}

// UnsubscribeTopics removes a list of topics from the list of topvs that the
// consumer is consuming.
func (c *Consumer) UnsubscribeTopics(topics []string) error {
	var err error
	for _, topic := range topics {
		for i, ctopic := range c.topics {
			if topic == ctopic {
				c.topics = append(c.topics[:i], c.topics[i+1:]...)
			}
			if c.subscriptions[topic] != nil {
				err = c.subscriptions[topic].Cancel()
				delete(c.subscriptions, topic)
			}

		}
	}

	return err
}

// ReadMessage will wait until there is a message from any one of the subscribed
// topics and return that message.
func (c *Consumer) ReadMessage() (*Message, error) {
	event := <-c.events

	switch e := event.(type) {
	case *Message:
		return e, nil
	case *Error:
		return nil, errors.New(e.reason)
	default:
		// ignore everything else
	}

	return nil, nil
}

// Poll returns the most event from all subscribed topics.
func (c *Consumer) Poll() Event {
	return <-c.events
}

// Start will start the consumption of all messages from the topics that the
// consumer is subscribed to. This needs to be called before attempting to read
// any messages.
func (c *Consumer) Start() {
	go c.consumeAllTopics()
}

// consumeAllTopics is an internal function that starts the consumption of messages
// from all topics that the consumer is subscribed to.
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

// consumeTopic is a helper function to pull the next message in a IPFS pubsub
// subscription and place it into our event channel to be processed either manually
// or by the ReadMessage function.
func (c *Consumer) consumeTopic(subscription *ipfs.PubSubSubscription) {
	for {
		record, err := subscription.Next()
		if err != nil {
			continue
		}

		var msg Message
		err = json.Unmarshal(record.Data(), &msg)
		if err != nil {
			continue
		}

		c.events <- &msg
	}
}

// sliceContainsSliceElement determines if a slice of strings contains at least
// one element from another slice of strings.
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
