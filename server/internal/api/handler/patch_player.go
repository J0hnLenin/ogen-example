package handler

import (
	"context"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

func (h *PlayerHandler) PatchPlayer(ctx context.Context, req *PlayerPartial, params PatchPlayerParams) (PatchPlayerRes, error) {

	var updated *models.Player
	var err error

	if req.Name.IsSet() && req.Score.IsSet() {
		updated, err = h.service.UpdatePlayer(ctx, params.ID, req.Name.Value, req.Score.Value)
	} else if req.Name.IsSet() {
		updated, err = h.service.UpdatePlayerName(ctx, params.ID, req.Name.Value)
	} else if req.Score.IsSet() {
		updated, err = h.service.UpdatePlayerScore(ctx, params.ID, req.Score.Value)
	} else {
		return &PatchPlayerBadRequest{
			Code:    http.StatusBadRequest,
			Message: "no fields to update",
		}, nil
	}
	if err != nil {
		return &PatchPlayerNotFound{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}, nil
	}
	return &Player{
		ID:    updated.ID,
		Name:  updated.Name,
		Score: updated.Score,
	}, nil
}
