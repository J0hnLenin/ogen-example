package clientservice

import (
	"context"
	"log"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

func (s *ClientService) PatchPlayer(ctx context.Context, id int, patch *playersapi.PlayerPartial) error {
	params := playersapi.PatchPlayerParams{ID: id}
	resp, err := s.client.PatchPlayer(ctx, patch, params)
	if err != nil {
		log.Printf("[PATCH] Error: %v", err)
		return err
	}
	if player, ok := resp.(*playersapi.Player); ok {
		log.Printf("[PATCH] ID=%d, Name=%s, Score=%.2f", player.ID, player.Name, player.Score)
	} else {
		log.Printf("[PATCH] Unexpected response type: %T", resp)
	}
	return nil
}
