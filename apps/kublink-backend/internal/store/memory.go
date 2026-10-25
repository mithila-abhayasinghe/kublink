package store

import (
	"context"
	"errors"
	"kublink-backend/internal/types"
	"sync"
	"sync/atomic"
)

type MemoryStore struct {
	mu      sync.RWMutex
	data    map[string]types.URLs
	counter atomic.Int64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]types.URLs),
	}
}

func (r *MemoryStore) Get(ctx context.Context, id string) (types.URLs, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, ok := r.data[id]
	if ok {
		return url, nil
	}
	return url, errors.New("not found")
}

func (r *MemoryStore) Create(ctx context.Context, u types.URLs) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[u.Id] = u
	return nil
}
