package payment

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/guilh/event-system/payment-service/internal/handler"
	"github.com/guilh/event-system/payment-service/internal/kafka"
	"github.com/guilh/event-system/payment-service/internal/repository"
	"github.com/guilh/event-system/payment-service/internal/usecase"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize PostgreSQL repository
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	defer db.Close()
	repo := repository.NewPostgresRepository(db)

	// Initialize Kafka producer
	broker := os.Getenv("KAFKA_BROKERS")
	producer := kafka.NewKafkaProducer(broker)

	// Initialize use case
	paymentUseCase := usecase.NewPaymentUseCase(repo, producer)

	// Initialize HTTP handler
	httpHandler := handler.NewHTTPHandler(paymentUseCase)

	// Set up Gin router
	router := gin.Default()
	router.POST("/payments", httpHandler.ProcessPayment)
	router.GET("/payments/:id", httpHandler.GetPayment)

	// Start the server
	router.Run(":8081")
}
