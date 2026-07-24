package playersserviceapi

import (
	"context"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
)

func (h *PlayerServiceAPI) UpdatePlayer(ctx context.Context, req *playersapi.PlayerInput, params playersapi.UpdatePlayerParams) (playersapi.UpdatePlayerRes, error) {
	player, err := h.service.UpdatePlayer(ctx, params.ID, req.Name, req.Score)
	if err != nil {
		status := http.StatusBadRequest
		return &playersapi.UpdatePlayerBadRequest{
			Code:    status,
			Message: err.Error(),
		}, nil
	}
	return &playersapi.Player{
		ID:    player.ID,
		Name:  player.Name,
		Score: player.Score,
	}, nil
}
