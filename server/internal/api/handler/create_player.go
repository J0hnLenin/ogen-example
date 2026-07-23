package handler

import (
	"context"
	"net/http"
)

func (h *PlayerHandler) CreatePlayer(ctx context.Context, req *PlayerInput) (CreatePlayerRes, error) {
	player, err := h.service.CreatePlayer(ctx, req.Name, req.Score)
	if err != nil {
		return &Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}
	return &Player{
		ID:    player.ID,
		Name:  player.Name,
		Score: player.Score,
	}, nil
}
