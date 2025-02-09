package router

import (
	"context"

	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/db"
)

type Handler struct {
	db *db.DB
}

func (h *Handler) Close() error {
	err := h.db.Close(context.Background())
	if err != nil {
		ctxerr.Handle(ctxerr.QuickWrap(context.Background(), err))
	}
	return nil
}

func NewHandler(ctx context.Context) (Handler, error) {
	db, err := db.New(ctx)
	if err != nil {
		return Handler{}, ctxerr.QuickWrap(context.Background(), err)
	}
	return Handler{db: db}, nil
}
