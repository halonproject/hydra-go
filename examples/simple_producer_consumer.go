package main

import (
	"fmt"
	"math/rand"
	"time"

	hydra "github.com/halonproject/hydra-go"
)

var topics = []string{"topic_1", "topic_2"}

func main() {
	config := hydra.DefaultConfig()

	producer := hydra.NewProducer(config)
	consumer, err := hydra.NewConsumer(config)
	if err != nil {
		fmt.Println("Error creating consumer", err.Error())
	}

	done := make(chan bool, 1)

	producer.AddTopics(topics)
	consumer.SubscribeTopics(topics)

	consumer.Start()
	go consumeMessages(consumer, done)

	produceMessages(producer)

	time.Sleep(time.Second * 30)
	done <- true
}

func consumeMessages(consumer *hydra.Consumer, done chan bool) {
	fmt.Println("Consuming messages...")
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			// continue to read messages
		}
		message, err := consumer.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
		}

		decoded, err := message.Decode()
		switch decoded.(type) {
		case []byte:
			fmt.Printf("%+v\n", string(decoded.([]byte)))
		}
	}

	fmt.Println("Done consuming...")
}

func produceMessages(producer *hydra.Producer) {
	for i := 0; i < 10; i++ {
		index := rand.Intn(1)
		msg := hydra.NewMessage([]byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("test_message_%d", i)), nil)
		if index == 0 {
			index = rand.Intn(len(topics))
			err := producer.Produce(topics[index], msg)
			if err != nil {
				fmt.Println("Error producing message:", err.Error())
			}
		} else {
			err := producer.ProduceAll(msg)
			if err != nil {
				fmt.Println("Error producing message:", err.Error())
			}
		}

		fmt.Printf("Produced message: %+v\n", msg)
		time.Sleep(time.Millisecond * 500)
	}
}
