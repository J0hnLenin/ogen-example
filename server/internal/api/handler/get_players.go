package handler

import (
	"context"
)

func (h *PlayerHandler) GetPlayers(ctx context.Context) ([]Player, error) {
	players, err := h.service.GetPlayers(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]Player, len(players))
	for i, p := range players {
		result[i] = Player{
			ID:    p.ID,
			Name:  p.Name,
			Score: p.Score,
		}
	}
	return result, nil
}
