package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {
	//1. create a pulsar client.
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//2. create a producer.
	topicName := "newTopicName"
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	//3. create a consumer.
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topicName,
		SubscriptionName: "subName",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	//4. send a test message by producer.
	err = producer.Send(context.Background(), pulsar.ProducerMessage{
		Payload: []byte(fmt.Sprintf("test")),
		DeliverAfter: 10 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}


	ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
	msg, err := consumer.Receive(ctx)
	fmt.Println("msg ======", string(msg.Payload()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(msg.Payload()))
	err = consumer.Ack(msg)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("consumer receive the message successful!")
	canc()

	ctx, canc = context.WithTimeout(context.Background(), 5*time.Second)
	msg, err = consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(msg.Payload()))
	canc()

}
