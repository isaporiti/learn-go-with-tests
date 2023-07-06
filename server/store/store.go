package server

import "fmt"

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore(scores map[string]int) *InMemoryPlayerStore {
	if scores == nil {
		return &InMemoryPlayerStore{make(map[string]int)}
	}
	return &InMemoryPlayerStore{scores}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	score, exists := s.scores[name]
	if !exists {
		return 0, fmt.Errorf("player not found")
	}
	return score, nil
}
