package domain

import (
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type Payment struct {
	ID        string        `json:"id"`         // Unique identifier for the payment
	OrderID   string        `json:"order_id"`   // ID of the order associated with this payment
	Amount    float64       `json:"amount"`     // Amount to be paid
	Status    PaymentStatus `json:"status"`     // Current status of the payment (e.g., "pending", "completed", "failed")
	CreatedAt time.Time     `json:"created_at"` // Timestamp when the payment was created
}
