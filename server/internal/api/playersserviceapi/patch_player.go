package playersserviceapi

import (
	"context"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

func (h *PlayerServiceAPI) PatchPlayer(ctx context.Context, req *playersapi.PlayerPartial, params playersapi.PatchPlayerParams) (playersapi.PatchPlayerRes, error) {

	var updated *models.Player
	var err error

	if req.Name.IsSet() && req.Score.IsSet() {
		updated, err = h.service.UpdatePlayer(ctx, params.ID, req.Name.Value, req.Score.Value)
	} else if req.Name.IsSet() {
		updated, err = h.service.UpdatePlayerName(ctx, params.ID, req.Name.Value)
	} else if req.Score.IsSet() {
		updated, err = h.service.UpdatePlayerScore(ctx, params.ID, req.Score.Value)
	} else {
		return &playersapi.PatchPlayerBadRequest{
			Code:    http.StatusBadRequest,
			Message: "no fields to update",
		}, nil
	}
	if err != nil {
		return &playersapi.PatchPlayerNotFound{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}, nil
	}
	return &playersapi.Player{
		ID:    updated.ID,
		Name:  updated.Name,
		Score: updated.Score,
	}, nil
}
