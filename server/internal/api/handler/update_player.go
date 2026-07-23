package handler

import (
	"context"
	"net/http"
)

func (h *PlayerHandler) UpdatePlayer(ctx context.Context, req *PlayerInput, params UpdatePlayerParams) (UpdatePlayerRes, error) {
	player, err := h.service.UpdatePlayer(ctx, params.ID, req.Name, req.Score)
	if err != nil {
		status := http.StatusBadRequest
		return &UpdatePlayerBadRequest{
			Code:    status,
			Message: err.Error(),
		}, nil
	}
	return &Player{
		ID:    player.ID,
		Name:  player.Name,
		Score: player.Score,
	}, nil
}
