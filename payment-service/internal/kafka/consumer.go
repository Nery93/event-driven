package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/guilh/event-system/payment-service/internal/domain"
	"github.com/guilh/event-system/payment-service/internal/usecase"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	broker  string
	groupID string
}

func NewKafkaConsumer(broker, groupID string) *KafkaConsumer {
	return &KafkaConsumer{broker: broker, groupID: groupID}
}

func (c *KafkaConsumer) ConsumeOrderEvents(uc *usecase.PaymentUseCase) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{c.broker},
		Topic:   "orders.created",
		GroupID: c.groupID,
	})
	defer r.Close()

	log.Println("Payment consumer listening on orders.created...")

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

		log.Printf("[orders.created] Processando pagamento para order %s", order.ID)

		if err := uc.ProcessPaymentFromOrder(context.Background(), &order); err != nil {
			log.Printf("Erro ao processar pagamento: %v", err)
		}
	}
}
