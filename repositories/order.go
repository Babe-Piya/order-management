package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type OrderRepository interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)
	CommitTransaction(ctx context.Context, tx pgx.Tx) error
	RollbackTransaction(ctx context.Context, tx pgx.Tx) error
	GetOrderByID(ctx context.Context, id int64, tx pgx.Tx) (Order, error)
	GetOrdersByPagination(ctx context.Context, page Pagination, tx pgx.Tx) ([]Order, error)
}
type orderRepository struct {
	DB *pgx.Conn
}

func NewOrderRepository(db *pgx.Conn) OrderRepository {
	return &orderRepository{DB: db}
}

type Order struct {
	ID           int64
	CustomerName string
	TotalAmount  float64
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (repo *orderRepository) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	return repo.DB.Begin(ctx)
}

func (repo *orderRepository) CommitTransaction(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (repo *orderRepository) RollbackTransaction(ctx context.Context, tx pgx.Tx) error {
	return tx.Rollback(ctx)
}

func (repo *orderRepository) GetOrderByID(ctx context.Context, id int64, tx pgx.Tx) (Order, error) {
	if tx == nil {
		tx, _ = repo.DB.BeginTx(ctx, pgx.TxOptions{})
	}

	// TODO: join order_items
	row, err := tx.Query(ctx, "SELECT * FROM orders WHERE id = $1", id)
	if err != nil {
		return Order{}, err
	}

	defer row.Close()

	result, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Order])
	if err != nil {
		return Order{}, err
	}

	return result, nil
}

type Pagination struct {
	Limit  int
	Offset int
	Page   int
}

func (repo *orderRepository) GetOrdersByPagination(ctx context.Context, page Pagination, tx pgx.Tx) ([]Order, error) {
	if tx == nil {
		tx, _ = repo.DB.BeginTx(ctx, pgx.TxOptions{})
	}

	rows, err := tx.Query(ctx, "SELECT * FROM OFFSET = $1 LIMIT = $2", page.Limit, page.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[Order])
	if err != nil {
		return nil, err
	}

	return result, nil
}
