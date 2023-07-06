package server

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
	sync.Mutex
}

func NewInMemoryPlayerStore(scores map[string]int) *InMemoryPlayerStore {
	if scores == nil {
		return &InMemoryPlayerStore{scores: map[string]int{}}
	}
	return &InMemoryPlayerStore{scores: scores}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	score, exists := s.scores[name]
	if !exists {
		return 0, fmt.Errorf("player not found")
	}
	return score, nil
}

func (s *InMemoryPlayerStore) ScoreWin(name string) {
	s.Lock()
	s.scores[name] += 1
	s.Unlock()
}
