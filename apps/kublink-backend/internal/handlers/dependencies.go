package handlers

import (
	"context"
	"kublink-backend/internal/types"
	"log/slog"
	"sync/atomic"
)

type URLStore interface {
	Create(ctx context.Context, url types.URLs) error
	Get(ctx context.Context, id string) (types.URLs, error)
}

type URLHandler struct {
	Store   URLStore
	counter atomic.Int64
	BaseURL string
	Logger  *slog.Logger
}
