# hydra-go

[![CircleCI](https://circleci.com/gh/HalonProject/hydra-go.svg?style=svg)](https://circleci.com/gh/HalonProject/hydra-go)
[![Coverage Status](https://coveralls.io/repos/github/HalonProject/hydra-go/badge.svg)](https://coveralls.io/github/HalonProject/hydra-go)

Golang implementation of `hydra` which is a library for producing and consuming
messages from multiple subscribed topics on IPFS.

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

- [ ] Go docs (In Progress)
- [ ] Unit/Coverage Testing (In Progress)
- [ ] Create standards for message headers (In Progress)
- [ ] Create standards for encrypted messages
- [ ] Create simple usage examples

## LICENSE

GNU GPL
