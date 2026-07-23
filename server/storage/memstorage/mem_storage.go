package memstorage

import "github.com/J0hnLenin/ogen-example/server/models"

type MemoryStorage struct {
	playerMap map[int]*models.Player
	nextID    int
}

func newMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		playerMap: make(map[int]*models.Player),
		nextID:    1,
	}
}
