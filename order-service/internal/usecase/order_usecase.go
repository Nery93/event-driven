package usecase

// OrderUseCase holds business logic for order operations.
// Depends only on interfaces defined in interfaces.go, never on concrete implementations.

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/guilh/event-system/order-service/internal/domain"
)


type OrderUseCase struct {
	repo      Repository
	publisher EventPublisher
}

func NewOrderUseCase(repo Repository, publisher EventPublisher) *OrderUseCase {
	return &OrderUseCase{
		repo:      repo,
		publisher: publisher,
	}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	// passo 1 — validação
	if order.CustomerID == "" {
		return nil, errors.New("customer_id é obrigatório")
	}
	if len(order.Items) == 0 {
		return nil, errors.New("order deve ter pelo menos 1 item")
	}

	// passo 2 — gerar ID único
	order.ID = uuid.New().String()

	// passo 3 — estado inicial
	order.Status = domain.OrderStatusPending
	order.CreatedAt = time.Now()

	// passo 4 — guardar na base de dados
	if err := uc.repo.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	// passo 5 — publicar evento Kafka
	if err := uc.publisher.PublishOrderCreated(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (uc *OrderUseCase) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	
	order, err := uc.repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
