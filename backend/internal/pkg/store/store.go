package store

import "sync"

type Store struct {
	mu          sync.RWMutex
	LastMessage string
}

func (s *Store) Set(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.LastMessage = msg
}

func (s *Store) Get() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.LastMessage
}
