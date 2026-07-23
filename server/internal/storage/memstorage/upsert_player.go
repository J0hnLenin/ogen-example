package memstorage

import (
	"context"
	"errors"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

func (m *MemoryStorage) UpsertPlayer(ctx context.Context, player *models.Player) (*models.Player, error) {
	if player.ID == 0 {
		player.ID = m.nextID
		m.nextID++
	} else if m.playerMap[player.ID] == nil {
		return nil, errors.New("storage error")
	}
	m.playerMap[player.ID] = player
	return player, nil
}
