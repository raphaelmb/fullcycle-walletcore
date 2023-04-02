package kafka

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
