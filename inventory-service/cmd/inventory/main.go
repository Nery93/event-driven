package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilh/event-system/inventory-service/internal/handler"
	"github.com/guilh/event-system/inventory-service/internal/kafka"
	"github.com/guilh/event-system/inventory-service/internal/repository"
	"github.com/guilh/event-system/inventory-service/internal/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	repo := repository.NewPostgresRepository(db, rdb)
	broker := os.Getenv("KAFKA_BROKERS")
	producer := kafka.NewKafkaProducer(broker)
	inventoryUseCase := usecase.NewInventoryUseCase(repo, producer)

	consumer := kafka.NewKafkaConsumer(broker, "inventory-service")
	go consumer.ConsumeOrderEvents(inventoryUseCase)
	go consumer.ConsumePaymentEvents(inventoryUseCase)

	httpHandler := handler.NewHTTPHandler(inventoryUseCase)
	router := gin.Default()
	router.POST("/inventory", httpHandler.CreateInventory)
	router.GET("/inventory/:id", httpHandler.GetProduct)
	router.Run(":8082")
}
