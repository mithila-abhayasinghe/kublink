package handlers

import (
	"kublink-backend/internal/store"
	"log/slog"
	"sync/atomic"
)

type URLHandler struct {
	Store   *store.URLStore
	counter atomic.Int64
	BaseURL string
	Logger  *slog.Logger
}
