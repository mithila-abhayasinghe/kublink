package store

import (
	"kublink-backend/internal/types"
	"sync"
)

type URLStore struct {
	mu   sync.RWMutex
	Data map[string]types.URLs
}

func NewURLStore() *URLStore {
	return &URLStore{
		Data: make(map[string]types.URLs),
	}
}

func (r *URLStore) Get(k string) (types.URLs, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, exists := r.Data[k]
	return url, exists
}

func (r *URLStore) Set(k string, u types.URLs) {
	r.mu.Lock()
	defer r.mu.RUnlock()
	r.Data[k] = u
}

func (r *URLStore) Delete(k string) {
	r.mu.Lock()
	defer r.mu.RUnlock()
	delete(r.Data, k)
}
