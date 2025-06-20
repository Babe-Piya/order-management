package service

import (
	"context"
	"log/slog"

	"github/Babe-piya/order-management/appconstant"
)

type UpdateStatusByIDRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type UpdateStatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (srv *orderService) UpdateStatusByID(ctx context.Context, req UpdateStatusByIDRequest) (*UpdateStatusResponse, error) {
	tx, err := srv.OrderRepo.BeginTransaction(ctx)
	if err != nil {
		slog.Error("begin transaction error:", err)

		return nil, err
	}

	defer func() {
		err = srv.OrderRepo.RollbackTransaction(ctx, tx)
		if err != nil {
			slog.Warn("rollback transaction error:", err)
		}
	}()

	err = srv.OrderRepo.UpdateStatusByID(ctx, req.Status, req.ID, tx)
	if err != nil {
		slog.Error("UpdateStatusByID error:", err)

		return nil, err
	}

	err = srv.OrderRepo.CommitTransaction(ctx, tx)
	if err != nil {
		slog.Error("commit transaction error:", err)

		return nil, err
	}

	return &UpdateStatusResponse{
		Code:    appconstant.SuccessCode,
		Message: appconstant.SuccessMsg,
	}, nil
}
