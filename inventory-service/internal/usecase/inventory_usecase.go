package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/guilh/event-system/inventory-service/internal/domain"
)

type InventoryUseCase struct {
	repo Repository
	pub  EventPublisher
}

func NewInventoryUseCase(repo Repository, pub EventPublisher) *InventoryUseCase {
	return &InventoryUseCase{
		repo: repo,
		pub:  pub,
	}
}

func (iu *InventoryUseCase) UpdateStock(ctx context.Context, inventory *domain.Inventory) error {

	// Se o stock for negativo, devo retornar um erro
	if inventory.Quantity < 0 {
		return errors.New("quantity deve ser maior ou igual a zero")
	}
	err := iu.repo.UpdateInventory(ctx, inventory)
	if err != nil {
		return err
	}

	err = iu.pub.PublishInventoryUpdated(ctx, inventory)
	if err != nil {
		return err
	}

	return nil
}

func (iu *InventoryUseCase) GetProductByID(ctx context.Context, id string) (*domain.Inventory, error) {

	inventory, err := iu.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (iu *InventoryUseCase) CreateInventory(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error) {
	
	if inventory.ProductID == "" {
		return nil, errors.New("product_id é obrigatório")
	}
	if inventory.Quantity < 0 {
		return nil, errors.New("quantity deve ser maior ou igual a zero")
	}
	inventory.ID = uuid.New().String()
	inventory.CreatedAt = time.Now()
	inventory.UpdatedAt = time.Now()

	if err := iu.repo.CreateInventory(ctx, inventory); err != nil {
		return nil, err
	}

	if err := iu.pub.PublishInventoryUpdated(ctx, inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}
