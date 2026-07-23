package memstorage

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/models"
)

func (m *MemoryStorage) GetPlayers(ctx context.Context) ([]*models.Player, error) {
	players := make([]*models.Player, 0, len(m.playerMap))
	for _, player := range m.playerMap {
		if player == nil {
			continue
		}
		players = append(players, player)
	}
	return players, nil
}
