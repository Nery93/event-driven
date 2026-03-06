package domain

import (
	"time"
)

type Notification struct {
	ID        string    `json:"id"`         // Unique identifier for the notification
	OrderID   string    `json:"order_id"`   // ID of the order associated with this notification
	Message   string    `json:"message"`    // Content of the notification
	CreatedAt time.Time `json:"created_at"` // Timestamp when the notification was created
}
