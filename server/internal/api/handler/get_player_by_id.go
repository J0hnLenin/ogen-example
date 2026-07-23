package handler

import (
	"context"
	"net/http"
)

func (h *PlayerHandler) GetPlayerById(ctx context.Context, params GetPlayerByIdParams) (GetPlayerByIdRes, error) {
	player, err := h.service.GetPlayerByID(ctx, params.ID)
	if err != nil {
		return &Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}, nil
	}
	return &Player{
		ID:    player.ID,
		Name:  player.Name,
		Score: player.Score,
	}, nil
}
