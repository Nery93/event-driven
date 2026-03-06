package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/guilh/event-system/inventory-service/internal/domain"
	"github.com/guilh/event-system/inventory-service/internal/usecase"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	broker  string
	groupID string
}

func NewKafkaConsumer(broker, groupID string) *KafkaConsumer {
	return &KafkaConsumer{broker: broker, groupID: groupID}
}

func (c *KafkaConsumer) ConsumeOrderEvents(uc *usecase.InventoryUseCase) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.broker},
		Topic:   "orders.created",
		GroupID: c.groupID,
	})
	defer r.Close()

	log.Println("Inventory consumer listening on orders.created...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			continue
		}

		var order domain.OrderEvent
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Printf("Erro ao deserializar: %v", err)
			continue
		}

		log.Printf("[orders.created] Atualizando inventário para order %s", order.ID)

		if err := uc.UpdateInventoryFromOrder(context.Background(), &order); err != nil {
			log.Printf("Erro ao atualizar inventário: %v", err)
		}
	}
}

func (c *KafkaConsumer) ConsumePaymentEvents(uc *usecase.InventoryUseCase) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.broker},
		Topic:   "payments.processed",
		GroupID: c.groupID,
	})
	defer r.Close()

	log.Println("Inventory consumer listening on payments.processed...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			continue
		}

		var payment domain.PaymentEvent
		if err := json.Unmarshal(msg.Value, &payment); err != nil {
			log.Printf("Erro ao deserializar payment event: %v", err)
			continue
		}

		if err := uc.HandlePaymentProcessed(context.Background(), &payment); err != nil {
			log.Printf("Erro ao processar payment event: %v", err)
		}
	}
}
