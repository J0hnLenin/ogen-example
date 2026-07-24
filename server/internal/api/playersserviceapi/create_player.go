package playersserviceapi

import (
	"context"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
)

func (h *PlayerServiceAPI) CreatePlayer(ctx context.Context, req *playersapi.PlayerInput) (playersapi.CreatePlayerRes, error) {
	player, err := h.service.CreatePlayer(ctx, req.Name, req.Score)
	if err != nil {
		return &playersapi.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}
	return &playersapi.Player{
		ID:    player.ID,
		Name:  player.Name,
		Score: player.Score,
	}, nil
}
