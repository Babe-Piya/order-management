package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	BeginTransaction(ctx context.Context) (pgx.Tx, error)
	CommitTransaction(ctx context.Context, tx pgx.Tx) error
	RollbackTransaction(ctx context.Context, tx pgx.Tx) error
	GetOrderByID(ctx context.Context, id int64, tx pgx.Tx) (*Order, error)
	GetOrdersByPagination(ctx context.Context, page int, rowOfPage int, tx pgx.Tx) ([]Order, error)
	GetCountOrder(ctx context.Context, tx pgx.Tx) (int, error)
	UpdateStatusByID(ctx context.Context, status string, id int64, tx pgx.Tx) error
	CreateOrder(ctx context.Context, data Order, tx pgx.Tx) (int64, error)
	CreateOrderItem(ctx context.Context, data []OrderItem, orderID int64, tx pgx.Tx) error
}
type orderRepository struct {
	DB *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{DB: db}
}

type Order struct {
	ID           int64
	CustomerName string
	TotalAmount  float64
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	OrderItems   []OrderItem
}

type OrderItem struct {
	ID          int64
	OrderID     int64
	ProductName string
	Quantity    int
	Price       float64
}

type OrderWithOrderItem struct {
	ID           int64
	CustomerName string
	TotalAmount  float64
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	OrderItemID  *int64
	ProductName  string
	Quantity     int
	Price        float64
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
