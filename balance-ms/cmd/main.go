package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *kafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *kafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume() error {
	consumer, err := kafka.NewConsumer(c.ConfigMap)
	if err != nil {
		fmt.Println("Consumer error", err.Error())
	}

	consumer.SubscribeTopics(c.Topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}
}

func main() {
	topics := []string{"transactions"}
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"client.id":         "consumer",
		"auto.offset.reset": "earliest",
	}
	consumer := NewConsumer(configMap, topics)

	consumer.Consume()
}
