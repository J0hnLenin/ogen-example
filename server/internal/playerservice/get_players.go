package playerservice

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

func (p *PlayerService) GetPlayers(ctx context.Context) ([]*models.Player, error) {
	return p.storage.GetPlayers(ctx)
}
