package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "foo",
		"auto.offset.reset": "earliest"})
	if err != nil {
		fmt.Errorf("%v", err)
	}

	topics := []string{"transaction"}
	err = consumer.SubscribeTopics(topics, nil)
	run := true
	fmt.Println("2")

	for run == true {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println(string(e.Value))
		case kafka.Error:

			fmt.Println(e.Error())
			run = false
		default:
		}
	}

	fmt.Println("1")
	consumer.Close()
}
