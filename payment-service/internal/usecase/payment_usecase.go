package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/guilh/event-system/payment-service/internal/domain"
)

type PaymentUseCase struct {
	repo      Repository
	publisher EventPublisher
}

func NewPaymentUseCase(repo Repository, publisher EventPublisher) *PaymentUseCase {
	return &PaymentUseCase{
		repo:      repo,
		publisher: publisher,
	}
}

func (pu *PaymentUseCase) ProcessPayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {

	if payment.OrderID == "" {
		return nil, errors.New("order_id é obrigatório")
	}

	payment.ID = uuid.New().String()
	payment.Status = domain.PaymentStatusPending

	if payment.Amount <= 0 {

		payment.Status = domain.PaymentStatusFailed
		err := pu.publisher.PublishPaymentFailed(ctx, payment)
		if err != nil {
			return nil, err
		}

		return nil, errors.New("amount deve ser maior que zero")
	}

	payment.Status = domain.PaymentStatusCompleted
	payment.CreatedAt = time.Now()

	err := pu.repo.CreatePayment(ctx, payment)
	if err != nil {
		return nil, err
	}

	err = pu.publisher.PublishPaymentProcessed(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (pu *PaymentUseCase) ProcessPaymentFromOrder(ctx context.Context, order *domain.OrderEvent) error {
	payment := &domain.Payment{
		OrderID: order.ID,
		Amount:  order.TotalAmount,
	}
	_, err := pu.ProcessPayment(ctx, payment)
	return err
}

func (pu *PaymentUseCase) GetPaymentByID(ctx context.Context, id string) (*domain.Payment, error) {
	payment, err := pu.repo.GetPaymentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
