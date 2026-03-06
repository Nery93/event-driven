package repository

// PostgresRepository implements the Repository interface defined in usecase/interfaces.go.

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/guilh/event-system/order-service/internal/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	query := `INSERT INTO orders (id, customer_id, items, total_amount, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.db.ExecContext(ctx, query, order.ID, order.CustomerID, itemsJSON, order.TotalAmount, order.Status, order.CreatedAt)
	return err
}

func (r *PostgresRepository) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {

	row := r.db.QueryRowContext(ctx, `SELECT id, customer_id, items, total_amount, status, created_at FROM orders WHERE id = $1`, id)

	var order domain.Order
	var itemsJSON []byte 

	err := row.Scan(&order.ID, &order.CustomerID, &itemsJSON, &order.TotalAmount, &order.Status, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	
	if err := json.Unmarshal(itemsJSON, &order.Items); err != nil {
		return nil, err
	}

	return &order, nil
}
