package kafka

// KafkaProducer implements the EventPublisher interface defined in usecase/interfaces.go.
// Publishes to topics: orders.created, orders.cancelled
import (
	"context"
	"encoding/json"

	"github.com/guilh/event-system/order-service/internal/domain"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writerCreated  *kafkago.Writer
	writerCanceled *kafkago.Writer
}

func NewKafkaProducer(broker string) *KafkaProducer {
	return &KafkaProducer{
		writerCreated: &kafkago.Writer{
			Addr:  kafkago.TCP(broker),
			Topic: "orders.created",
		},
		writerCanceled: &kafkago.Writer{
			Addr:  kafkago.TCP(broker),
			Topic: "orders.cancelled",
		},
	}
}

func (p *KafkaProducer) PublishOrderCreated(ctx context.Context, order *domain.Order) error {

	bytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = p.writerCreated.WriteMessages(ctx, kafkago.Message{
		Value: bytes,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProducer) PublishOrderCanceled(ctx context.Context, order *domain.Order) error {
	
	bytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = p.writerCanceled.WriteMessages(ctx, kafkago.Message{
		Value: bytes,
	})
	if err != nil {
		return err
	}

	return nil
}
