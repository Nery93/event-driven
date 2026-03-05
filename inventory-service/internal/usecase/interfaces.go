package usecase

import (
	"context"

	"github.com/guilh/event-system/inventory-service/internal/domain"
)

type Repository interface {
	UpdateInventory(ctx context.Context, inventory *domain.Inventory) error
	GetProductByID(ctx context.Context, id string) (*domain.Inventory, error)
	CreateInventory(ctx context.Context, inventory *domain.Inventory) error
}

type EventPublisher interface {
	PublishInventoryUpdated(ctx context.Context, inventory *domain.Inventory) error
}
