package hydra

import (
	"fmt"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type Producer struct {
	topics []string
	conn   *ipfs.Shell
}

func NewProducer(config *Config) *Producer {
	dsn := fmt.Sprintf("http://%s:%s", config.IPFSAddr, config.IPFSPort)
	producer := &Producer{
		topics: config.Topics,
		conn:   ipfs.NewShell(dsn),
	}
	return producer
}

func (p *Producer) Topics() []string {
	return p.topics
}

func (p *Producer) AddTopic(topic string) {
	p.AddTopics([]string{topic})
}

func (p *Producer) AddTopics(topics []string) {
	for _, topic := range topics {
		if !sliceContainsString(p.topics, topic) {
			p.topics = append(p.topics, topic)
		}
	}
}

func (p *Producer) RemoveTopic(topic string) {
	p.RemoveTopics([]string{topic})
}

func (p *Producer) RemoveTopics(topics []string) {
	for _, topic := range topics {
		for i, subscription := range p.topics {
			if topic == subscription {
				p.topics = append(p.topics[:i], p.topics[i+1:]...)
			}
		}
	}
}

func (p *Producer) Produce(topic string, msg *Message) error {
	if !sliceContainsString(p.topics, topic) {
		return fmt.Errorf("Cannot publish message to unsubscribed topic \"%s\"", topic)
	}

	msgBytes, err := msg.Encode()
	if err != nil {
		return err
	}

	err = p.conn.PubSubPublish(topic, string(msgBytes))

	return err
}

func (p *Producer) ProduceAll(msg *Message) error {
	msgBytes, err := msg.Encode()
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

func sliceContainsString(haystack []string, needle string) bool {
	for _, stack := range haystack {
		if stack == needle {
			return true
		}
	}

	return false
}
