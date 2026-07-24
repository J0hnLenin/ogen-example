package clientservice

import (
	"context"
	"log"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

func (s *ClientService) CreatePlayer(ctx context.Context, input *playersapi.PlayerInput) error {
	resp, err := s.client.CreatePlayer(ctx, input)
	if err != nil {
		log.Printf("[CREATE] Error: %v", err)
		return err
	}
	if player, ok := resp.(*playersapi.Player); ok {
		log.Printf("[CREATE] ID=%d, Name=%s, Score=%.2f", player.ID, player.Name, player.Score)
	} else {
		log.Printf("[CREATE] Unexpected response type: %T", resp)
	}
	return nil
}
