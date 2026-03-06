package kafka

import (
	"context"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	readerPaymentsProcessed *kafkago.Reader
	readerPaymentsFailed    *kafkago.Reader
	readerInventoryUpdated  *kafkago.Reader
}

func NewKafkaConsumer(broker string) *KafkaConsumer {
	return &KafkaConsumer{
		readerPaymentsProcessed: kafkago.NewReader(kafkago.ReaderConfig{
			Brokers: []string{broker},
			Topic:   "payments.processed",
			GroupID: "notification-service",
		}),
		readerPaymentsFailed: kafkago.NewReader(kafkago.ReaderConfig{
			Brokers: []string{broker},
			Topic:   "payments.failed",
			GroupID: "notification-service",
		}),
		readerInventoryUpdated: kafkago.NewReader(kafkago.ReaderConfig{
			Brokers: []string{broker},
			Topic:   "inventory.updated",
			GroupID: "notification-service",
		}),
	}
}

func (c *KafkaConsumer) Start(ctx context.Context) {
	go c.consumePaymentsProcessed(ctx)
	go c.consumePaymentsFailed(ctx)
	go c.consumeInventoryUpdated(ctx)
}

func (c *KafkaConsumer) consumePaymentsProcessed(ctx context.Context) {
	for {
		msg, err := c.readerPaymentsProcessed.ReadMessage(ctx)
		if err != nil {
			continue // em aplicação real, considero um DeadLetterQueue ou algo do tipo para lidar com erros de leitura
		}
		log.Printf("[payments.processed] %s", string(msg.Value))
	}
}

func (c *KafkaConsumer) consumePaymentsFailed(ctx context.Context) {
	for {
		msg, err := c.readerPaymentsFailed.ReadMessage(ctx)
		if err != nil {
			continue // em aplicação real, considero um DeadLetterQueue ou algo do tipo para lidar com erros de leitura
		}
		log.Printf("[payments.failed] %s", string(msg.Value))
	}
}

func (c *KafkaConsumer) consumeInventoryUpdated(ctx context.Context) {
	for {
		msg, err := c.readerInventoryUpdated.ReadMessage(ctx)
		if err != nil {
			continue // em aplicação real, considero um DeadLetterQueue ou algo do tipo para lidar com erros de leitura
		}
		log.Printf("[inventory.updated] %s", string(msg.Value))
	}
}

func (c *KafkaConsumer) Close() error {
	if err := c.readerPaymentsProcessed.Close(); err != nil {
		return err
	}
	if err := c.readerPaymentsFailed.Close(); err != nil {
		return err
	}
	if err := c.readerInventoryUpdated.Close(); err != nil {
		return err
	}
	return nil
}

