package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/guilh/event-system/inventory-service/internal/domain"
	"github.com/redis/go-redis/v9"
)

type PostgresRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewPostgresRepository(db *sql.DB, rdb *redis.Client) *PostgresRepository {
	return &PostgresRepository{
		db:  db,
		rdb: rdb,
	}
}

// Implementação dos métodos do repositório usando PostgreSQL e Redis
func (r *PostgresRepository) CreateInventory(ctx context.Context, inventory *domain.Inventory) error {

	query := `INSERT INTO inventory (id, product_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, inventory.ID, inventory.ProductID, inventory.Quantity, inventory.CreatedAt, inventory.UpdatedAt)
	return err
}

func (r *PostgresRepository) GetProductByID(ctx context.Context, id string) (*domain.Inventory, error) {

	val, err := r.rdb.Get(ctx, id).Result()
	if err == nil {
		var inventory domain.Inventory
		err = json.Unmarshal([]byte(val), &inventory)
		return &inventory, nil
	}

	row := r.db.QueryRowContext(ctx, `SELECT id, product_id, quantity, created_at, updated_at FROM inventory WHERE product_id = $1`, id)
	var inventory domain.Inventory
	if err := row.Scan(&inventory.ID, &inventory.ProductID, &inventory.Quantity, &inventory.CreatedAt, &inventory.UpdatedAt); err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(inventory)
	r.rdb.Set(ctx, id, bytes, time.Minute*10)

	return &inventory, nil
}

func (r *PostgresRepository) UpdateInventory(ctx context.Context, inventory *domain.Inventory) error {

	query := `UPDATE inventory SET quantity = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, inventory.Quantity, inventory.UpdatedAt, inventory.ID)
	return err
	
}