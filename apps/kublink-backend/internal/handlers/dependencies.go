package handlers

import (
	"kublink-backend/internal/store"
	"sync/atomic"
)

type URLHandler struct {
	Store   *store.URLStore
	counter atomic.Int64
	BaseURL string
}
