package main

import (
	"context"
	"log"
	"os"

	"github.com/guilh/event-system/notification-service/internal/kafka"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	broker := os.Getenv("KAFKA_BROKERS")
	consumer := kafka.NewKafkaConsumer(broker)

	ctx := context.Background()
	consumer.Start(ctx)

	log.Println("Notification Service is running...")
	select {} // mantém o serviço rodando
}
