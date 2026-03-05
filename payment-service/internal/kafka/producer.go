package kafka

import (
	"context"
	"encoding/json"

	"github.com/guilh/event-system/payment-service/internal/domain"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writerProcessed *kafkago.Writer
	writerFailed    *kafkago.Writer
}

func NewKafkaProducer(broker string) *KafkaProducer {
	return &KafkaProducer{
		writerProcessed: &kafkago.Writer{
			Addr:  kafkago.TCP(broker),
			Topic: "payments.processed",
		},
		writerFailed: &kafkago.Writer{
			Addr:  kafkago.TCP(broker),
			Topic: "payments.failed",
		},
	}
}

func (p *KafkaProducer) PublishPaymentProcessed(ctx context.Context, payment *domain.Payment) error {
	bytes, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	err = p.writerProcessed.WriteMessages(ctx, kafkago.Message{
		Value: bytes,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProducer) PublishPaymentFailed(ctx context.Context, payment *domain.Payment) error {
	bytes, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	err = p.writerFailed.WriteMessages(ctx, kafkago.Message{
		Value: bytes,
	})
	if err != nil {
		return err
	}

	return nil
}
