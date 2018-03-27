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
	dsn := fmt.Sprintf("https://%s:%s", config.IPFSAddr, config.IPFSPort)
	producer := &Producer{
		topics: config.Topics,
		conn:   ipfs.NewShell(dsn),
	}
	return producer
}

func (p *Producer) AddTopic(topic string) {
	p.AddTopics([]string{topic})
}

func (p *Producer) AddTopics(topics []string) {
	for _, topic := range topics {
		p.topics = append(p.topics, topic)
	}
}

func (p *Producer) RemoveTopic(topic string) {
	p.RemoveTopics([]string{topic})
}

func (p *Producer) RemoveTopics(topics []string) {
	newTopics := make([]string, 0)
	for _, topic := range topics {
		for _, subscription := range p.topics {
			if topic != subscription {
				newTopics = append(newTopics, subscription)
			}
		}
	}

	p.topics = newTopics
}

func (p *Producer) Produce(topic string, msg *Message) error {
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
