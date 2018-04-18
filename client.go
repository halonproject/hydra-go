package hydra

import ipfs "github.com/ipfs/go-ipfs-api"

type IPFSClient interface {
	PubSubPublish(topic, data string) error
	PubSubSubscribe(topic string) (*ipfs.PubSubSubscription, error)
}
