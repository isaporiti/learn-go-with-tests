package store

import (
	"fmt"
	"sync"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	ScoreWin(name string)
}

type InMemoryPlayerStore struct {
	scores map[string]int
	mutex  *sync.Mutex
}

type Option func(store *InMemoryPlayerStore) error

func WithScores(scores map[string]int) Option {
	return func(store *InMemoryPlayerStore) error {
		if scores == nil {
			return fmt.Errorf("scores can't be nil")
		}
		store.scores = scores
		return nil
	}
}

func WithMutex(mutex *sync.Mutex) Option {
	return func(store *InMemoryPlayerStore) error {
		if mutex == nil {
			return fmt.Errorf("mutex can't be nil")
		}
		store.mutex = mutex
		return nil
	}
}

func NewInMemoryPlayerStore(options ...Option) (*InMemoryPlayerStore, error) {
	store := InMemoryPlayerStore{
		scores: map[string]int{},
		mutex:  &sync.Mutex{},
	}
	for _, option := range options {
		err := option(&store)
		if err != nil {
			return nil, fmt.Errorf("could not create store: %s", err.Error())
		}
	}
	return &store, nil
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	score, exists := s.scores[name]
	if !exists {
		return 0, fmt.Errorf("player not found")
	}
	return score, nil
}

func (s *InMemoryPlayerStore) ScoreWin(name string) {
	s.mutex.Lock()
	s.scores[name] += 1
	s.mutex.Unlock()
}
