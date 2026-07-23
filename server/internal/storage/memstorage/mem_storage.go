package memstorage

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

type MemoryStorage struct {
	playerMap map[int]*models.Player
	nextID    int
}

func NewMemoryStorage(ctx context.Context) *MemoryStorage {
	return &MemoryStorage{
		playerMap: make(map[int]*models.Player),
		nextID:    1,
	}
}
