package playersserviceapi

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
)

func (h *PlayerServiceAPI) GetPlayers(ctx context.Context) ([]playersapi.Player, error) {
	players, err := h.service.GetPlayers(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]playersapi.Player, len(players))
	for i, p := range players {
		result[i] = playersapi.Player{
			ID:    p.ID,
			Name:  p.Name,
			Score: p.Score,
		}
	}
	return result, nil
}
