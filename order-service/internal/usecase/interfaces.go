package usecase

// Repository and EventPublisher interfaces are defined here.
// repository/ and kafka/ packages must implement these.

import (
	"context"

	"github.com/guilh/event-system/order-service/internal/domain"
)

// Repository defines the methods for interacting with the order data store.
type Repository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrderByID(ctx context.Context, id string) (*domain.Order, error)
}

// EventPublisher defines the methods for publishing order-related events to the message broker.
type EventPublisher interface {
	PublishOrderCreated(ctx context.Context, order *domain.Order) error
	PublishOrderCanceled(ctx context.Context, order *domain.Order) error
}
