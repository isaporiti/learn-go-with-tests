package store

import (
	"encoding/json"
	"fmt"
	"io"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func NewFileSystemStore(database io.ReadSeeker) *FileSystemStore {
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
