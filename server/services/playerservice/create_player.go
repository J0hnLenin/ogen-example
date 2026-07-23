package playerservice

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/models"
)

func (p *PlayerService) CreatePlayer(ctx context.Context, name string, score float64) (*models.Player, error) {
	player := &models.Player{
		ID:    0,
		Name:  name,
		Score: score,
	}

	return p.storage.UpsertPlayer(ctx, player)
}
