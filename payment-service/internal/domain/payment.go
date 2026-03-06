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


type OrderEvent struct {
	ID          string  `json:"id"`
	CustomerID  string  `json:"customer_id"`
	TotalAmount float64 `json:"total_amount"`
}

type Payment struct {
	ID        string        `json:"id"`         
	OrderID   string        `json:"order_id"`   
	Amount    float64       `json:"amount"`     
	Status    PaymentStatus `json:"status"`     
	CreatedAt time.Time     `json:"created_at"` 
