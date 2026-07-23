package memstorage

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/models"
)

func (m *MemoryStorage) GetPlayerByID(ctx context.Context, id int) (*models.Player, error) {
	return m.playerMap[id], nil
}
