# hydra-go

Golang implementation of `hydra` which is a library for producing and consuming
messages from multiple subscribed topics on IPFS. Currently on consumption is allowed
since the `go-ipfs-api` does not seems to support producing messages.

## Motivation

There exists many publish-subscribe/message queue system but IPFS has a built in
pubsub system. This is a decentralized system that we can exploit to build more
complex systems similar to Kafka, RabbitMQ. This project attempts to build some
standards around producing and consuming messages from multiple topics hosted on IPFS.

Those standards include:

- Application specific messages:
  - JSON
  - XML
- Encrypted messages
  - AES
  - Asymmetric Encryption
  - Symmetric Encryption

## TODO

- [ ] Go docs
- [ ] Unit/Coverage Testing
- [ ] Create standards for message headers
- [ ] Create standards for encrypted messages

## LICENSE

GNU GPL
