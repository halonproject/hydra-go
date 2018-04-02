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

## Testing

In order to do any testing you will have to have an IPFS client that you can
connect to that has pubsub enabled. You can run one locally by running the following
command:

```
$ ipfs daemon --enable-pubsub-experiment
```

In another terminal you can then run all tests:

```
$ go test
```

## Example usage

In the `examples` folder there is a simple go program that creates a producer and
consumer with the same topics. This shows how you can use this package to make a
simple project to read and write messages without having to set up any infrastructure
to provide the messaging service.

When you run the program you should see output similar to:

```
$ go run examples/simple_producer_consumer.go
Consuming messages...
Produced message: 1522692256: [0=test_message_0]
Consumed message: test_message_0
Consumed message: test_message_1
Produced message: 1522692256: [1=test_message_1]
Produced message: 1522692257: [2=test_message_2]
Consumed message: test_message_2
Produced message: 1522692257: [3=test_message_3]
Consumed message: test_message_3
Produced message: 1522692258: [4=test_message_4]
Consumed message: test_message_4
Produced message: 1522692258: [5=test_message_5]
Consumed message: test_message_5
Produced message: 1522692259: [6=test_message_6]
Consumed message: test_message_6
Produced message: 1522692259: [7=test_message_7]
Consumed message: test_message_7
Produced message: 1522692260: [8=test_message_8]
Consumed message: test_message_8
Produced message: 1522692260: [9=test_message_9]
Consumed message: test_message_9
```

## Contributing

Feel free to check out the `CONTRIBUTING.md` for the guidelines on contributing
to the project.  

## TODO

- [ ] Create standards for message headers (In Progress)
- [ ] Create standards for encrypted messages

## LICENSE

GPL
