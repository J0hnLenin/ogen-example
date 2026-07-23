package playerservice

import (
	"context"
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
)

func (p *PlayerService) GetPlayerByID(ctx context.Context, id int) (*models.Player, error) {
	player, err := p.storage.GetPlayerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, errors.New("player not found")
	}
	return player, nil
}
