package usecase

import (
	"context"

	"github.com/guilh/event-system/payment-service/internal/domain"
)

type Repository interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) error
	GetPaymentByID(ctx context.Context, id string) (*domain.Payment, error)
}

type EventPublisher interface {
	PublishPaymentProcessed(ctx context.Context, payment *domain.Payment) error
	PublishPaymentFailed(ctx context.Context, payment *domain.Payment) error
}