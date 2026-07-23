package playerservice

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/models"
)

func (p *PlayerService) UpdatePlayerName(ctx context.Context, id int, newName string) (*models.Player, error) {
	return nil, nil
}
