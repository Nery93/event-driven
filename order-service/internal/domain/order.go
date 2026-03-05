package domain

// Order represents a customer's order in the system.

import (
	"time"
)

type Order struct {
	ID          string             `json:"id"`           // Unique identifier for the order
	CustomerID  string             `json:"customer_id"`  // ID of the customer who placed the order
	Items       []Item             `json:"items"`        // List of items in the order
	TotalAmount float64            `json:"total_amount"` // Total amount for the order
	Status      OrderStatus        `json:"status"`       // Current status of the order (e.g., "pending", "shipped", "delivered")
	CreatedAt   time.Time          `json:"created_at"`   // Timestamp when the order was created
}

// Item represents a single item in an order.

type Item struct {
	ProductID string  `json:"product_id"` // ID of the product
	Quantity  int     `json:"quantity"`   // Quantity of the product ordered
	Price     float64 `json:"price"`      // Price of a single unit of the product
}
