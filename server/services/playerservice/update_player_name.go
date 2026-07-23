package playerservice

import (
	"context"
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
)

func (p *PlayerService) UpdatePlayerName(ctx context.Context, id int, newName string) (*models.Player, error) {
	player, err := p.storage.GetPlayerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, errors.New("player not found")
	}

	player.Name = newName
	player, err = p.storage.UpsertPlayer(ctx, player)

	if err != nil {
		return nil, err
	}
	return player, nil
}
