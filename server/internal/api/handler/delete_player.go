package handler

import (
	"context"
	"net/http"
)

func (h *PlayerHandler) DeletePlayer(ctx context.Context, params DeletePlayerParams) (DeletePlayerRes, error) {
	err := h.service.DeletePlayer(ctx, params.ID)
	if err != nil {
		return &Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}, nil
	}
	return &DeletePlayerNoContent{}, nil
}
