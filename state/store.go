package state

import (
	"sync"
)

// Store represents a global state store
type Store struct {
	values map[string]any
	mu     sync.RWMutex
}

// NewStore creates a new state store
func NewStore() *Store {
	return &Store{
		values: make(map[string]interface{}),
	}
}

// Set sets a value in the store
func (s *Store) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = value
}

// Get retrieves a value from the store
func (s *Store) Get(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.values[key]
}

// GetInt gets an int value with a default if not found
func (s *Store) GetInt(key string, defaultValue int) int {
	value := s.Get(key)
	if value == nil {
		return defaultValue
	}
	if v, ok := value.(int); ok {
		return v
	}
	return defaultValue
}

// Global store instance
var GlobalStore = NewStore()
