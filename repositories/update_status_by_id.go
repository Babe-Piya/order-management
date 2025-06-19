package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (repo *orderRepository) UpdateStatusByID(ctx context.Context, status string, id int64, tx pgx.Tx) error {
	if tx == nil {
		tx, _ = repo.DB.Begin(ctx)
	}

	row, err := tx.Exec(ctx, "UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3", status, time.Now(), id)
	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}
