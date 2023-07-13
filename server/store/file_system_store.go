package store

import (
	"encoding/json"
	"fmt"
	"io"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
	return &FileSystemStore{database: database}
}

func (s *FileSystemStore) GetLeague() ([]server.Player, error) {
	s.database.Seek(0, 0)
	var league []server.Player
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
	for _, player := range league {
		if player.Name == name {
			return player.Wins, nil
		}
	}
	return 0, fmt.Errorf("could find player '%s'", name)
}

func (s *FileSystemStore) ScoreWin(name string) error {
	league, err := s.GetLeague()
	if err != nil {
		return fmt.Errorf("could not score win for player '%s'", name)
	}
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}
	s.database.Seek(0, 0)
	json.NewEncoder(s.database).Encode(league)
	return nil
}
