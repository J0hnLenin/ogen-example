package playersserviceapi

import (
	"context"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
)

func (h *PlayerServiceAPI) GetPlayerById(ctx context.Context, params playersapi.GetPlayerByIdParams) (playersapi.GetPlayerByIdRes, error) {
	player, err := h.service.GetPlayerByID(ctx, params.ID)
	if err != nil {
		return &playersapi.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}, nil
	}
	return &playersapi.Player{
		ID:    player.ID,
		Name:  player.Name,
		Score: player.Score,
	}, nil
}
