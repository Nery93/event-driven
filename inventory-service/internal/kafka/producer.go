package kafka

import (
	"context"
	"encoding/json"

	"github.com/guilh/event-system/inventory-service/internal/domain"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	inventoryUpdatedWriter *kafkago.Writer
}

func NewKafkaProducer(broker string) *KafkaProducer {
	return &KafkaProducer{
		inventoryUpdatedWriter: &kafkago.Writer{
			Addr:  kafkago.TCP(broker),
			Topic: "inventory.updated",
		},
	}
}

func (p *KafkaProducer) PublishInventoryUpdated(ctx context.Context, inventory *domain.Inventory) error {
	
	bytes, err := json.Marshal(inventory)
	if err != nil {
		return err
	}

	return p.inventoryUpdatedWriter.WriteMessages(ctx, kafkago.Message{
		Key:   []byte(inventory.ID),
		Value: bytes,
	})
}