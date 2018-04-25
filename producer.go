package hydra

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	message "github.com/halonproject/hydra-proto-go"
	ipfs "github.com/ipfs/go-ipfs-api"
)

// Producer is a high level message producer that publishes messages to a single
// or multiple topics on IPFS pubsub.
type Producer struct {
	topics []string
	conn   *ipfs.Shell
}

// NewProducer creates a new producer connected to a IPFS client specified in the
// configuration.
func NewProducer(config *Config) *Producer {
	dsn := fmt.Sprintf("http://%s:%s", config.IPFSAddr, config.IPFSPort)
	producer := &Producer{
		topics: config.Topics,
		conn:   ipfs.NewShell(dsn),
	}
	return producer
}

// Topics returns a list of all topics that the producer is subscibed to.
func (p *Producer) Topics() []string {
	return p.topics
}

// AddTopic will add a topic to the list of topics the producer will publish
// messages to.
func (p *Producer) AddTopic(topic string) {
	p.AddTopics([]string{topic})
}

// AddTopics will add a list of topics to the list of topics the producer will
// publish messages to.
func (p *Producer) AddTopics(topics []string) {
	for _, topic := range topics {
		if !sliceContainsString(p.topics, topic) {
			p.topics = append(p.topics, topic)
		}
	}
}

// RemoveTopic will will remove a topic from the producers list of topics.
func (p *Producer) RemoveTopic(topic string) {
	p.RemoveTopics([]string{topic})
}

// RemoveTopics removes a list of topics from the producers list of topics.
func (p *Producer) RemoveTopics(topics []string) {
	for _, topic := range topics {
		for i, subscription := range p.topics {
			if topic == subscription {
				p.topics = append(p.topics[:i], p.topics[i+1:]...)
			}
		}
	}
}

// Produce will publish a message to a specific topic on IPFS. If the topic provided
// is not in producers list of subscribed topics it will throw an error.
func (p *Producer) Produce(topic string, msg *message.Message) error {
	if !sliceContainsString(p.topics, topic) {
		return fmt.Errorf("Cannot publish message to unsubscribed topic \"%s\"", topic)
	}

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	err = p.conn.PubSubPublish(topic, string(msgBytes))

	return err
}

// ProduceAll will publish a message to all of the topics that the producer is
// subscribed to.
func (p *Producer) ProduceAll(msg *message.Message) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	for _, topic := range p.topics {
		err = p.conn.PubSubPublish(topic, string(msgBytes))
		if err != nil {
			return err
		}
	}

	return nil
}

// sliceContainsString is a helper function to determine if a slice of string
// contains a specified string.
func sliceContainsString(haystack []string, needle string) bool {
	for _, stack := range haystack {
		if stack == needle {
			return true
		}
	}

	return false
}
