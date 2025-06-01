package exsync

import (
	"sync"
)

type LockMap[T comparable] struct {
	mu    sync.Mutex
	locks map[T]*sync.Mutex
}

// NewLockMap creates a new instance of LockMap.
func NewLockMap[T comparable]() *LockMap[T] {
	return &LockMap[T]{locks: make(map[T]*sync.Mutex)}
}

// Lock acquires the mutex for the given key.
func (lm *LockMap[T]) Lock(key T) {
	lm.mu.Lock()
	m, exists := lm.locks[key]
	if !exists {
		m = &sync.Mutex{}
		lm.locks[key] = m
	}
	lm.mu.Unlock()

	m.Lock()
}

// Unlock releases the mutex for the given key.
func (lm *LockMap[T]) Unlock(key T) {
	lm.mu.Lock()
	m, exists := lm.locks[key]
	lm.mu.Unlock()

	if exists {
		m.Unlock()
	}
}
