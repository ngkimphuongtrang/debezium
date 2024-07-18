package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/trangnkp/debezium/internal/config"
	"log"
)

type Consumer struct {
	*kafka.Consumer
}

func New(cfg *config.Config) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Host,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return &Consumer{
		Consumer: consumer,
	}, nil
}

func (c *Consumer) SubscribeTopic(topic string) error {
	err := c.Consumer.Subscribe(topic, nil)
	if err != nil {
		log.Println("subscribe error", err)
		return err
	}
	return nil
}

func (c *Consumer) ReadMessage() (*kafka.Message, error) {
	msg, err := c.Consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
