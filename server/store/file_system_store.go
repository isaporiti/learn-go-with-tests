package store

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/isaporiti/learn-go-with-tests/server/server"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
	return &FileSystemStore{database: database}
}

func (s *FileSystemStore) GetLeague() (server.League, error) {
	s.database.Seek(0, 0)
	var league server.League
	err := json.NewDecoder(s.database).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("couln't decode league: %v", err)
	}
	return league, nil
}

func (s *FileSystemStore) GetPlayerScore(name string) (int, error) {
	league, err := s.GetLeague()
	if err != nil {
		return 0, fmt.Errorf("could not get player '%s' score: %v", name, err)
	}
	player := league.Find(name)
	if player == nil {
		return 0, nil
	}
	return player.Wins, nil
}

func (s *FileSystemStore) ScoreWin(name string) error {
	league, err := s.GetLeague()
	if err != nil {
		return fmt.Errorf("could not get league: %v", err)
	}
	player := league.Find(name)
	if player == nil {
		player := server.Player{Name: name, Wins: 1}
		league = append(league, player)
	} else {
		player.Wins++
	}
	s.database.Seek(0, 0)
	json.NewEncoder(s.database).Encode(league)
	return nil
}
