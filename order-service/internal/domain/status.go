package domain

type OrderStatus string
const (
	// OrderStatusPending indicates that the order is pending and has not been processed yet.
	OrderStatusPending OrderStatus = "pending"
	// OrderStatusCanceled indicates that the order has been canceled.
	OrderStatusCanceled OrderStatus = "canceled"
)