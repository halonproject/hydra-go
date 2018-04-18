package hydra

import ipfs "github.com/ipfs/go-ipfs-api"

type mockIPFSClient struct{}

func newMockIPFSClient() *mockIPFSClient {
	return &mockIPFSClient{}
}

func (client *mockIPFSClient) PubSubPublish(topic, data string) error {
	return nil
}

func (client *mockIPFSClient) PubSubSubscribe(topic string) (*ipfs.PubSubSubscription, error) {
	return &ipfs.PubSubSubscription{}, nil
}
