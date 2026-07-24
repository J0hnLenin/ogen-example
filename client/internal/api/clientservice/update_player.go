package clientservice

import (
	"context"
	"log"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

func (s *ClientService) UpdatePlayer(ctx context.Context, id int, input *playersapi.PlayerInput) error {
	params := playersapi.UpdatePlayerParams{ID: id}
	resp, err := s.client.UpdatePlayer(ctx, input, params)
	if err != nil {
		log.Printf("[UPDATE] Error: %v", err)
		return err
	}
	if player, ok := resp.(*playersapi.Player); ok {
		log.Printf("[UPDATE] ID=%d, Name=%s, Score=%.2f", player.ID, player.Name, player.Score)
	} else {
		log.Printf("[UPDATE] Unexpected response type: %T", resp)
	}
	return nil
}
