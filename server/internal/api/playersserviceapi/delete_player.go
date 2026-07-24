package playersserviceapi

import (
	"context"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
)

func (h *PlayerServiceAPI) DeletePlayer(ctx context.Context, params playersapi.DeletePlayerParams) (playersapi.DeletePlayerRes, error) {
	err := h.service.DeletePlayer(ctx, params.ID)
	if err != nil {
		return &playersapi.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}, nil
	}
	return &playersapi.DeletePlayerNoContent{}, nil
}
