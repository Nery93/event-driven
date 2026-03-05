package repository


import (	"context"
	"database/sql"

	"github.com/guilh/event-system/payment-service/internal/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreatePayment(ctx context.Context, payment *domain.Payment) error {
	query := `INSERT INTO payments (id, order_id, amount, status, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, payment.ID, payment.OrderID, payment.Amount, payment.Status, payment.CreatedAt)
	return err
}

func (r *PostgresRepository) GetPaymentByID(ctx context.Context, id string) (*domain.Payment, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, order_id, amount, status, created_at FROM payments WHERE id = $1`, id)
	var payment domain.Payment
	err := row.Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Status, &payment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}