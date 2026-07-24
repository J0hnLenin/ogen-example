package clientservice

import (
	"context"
	"log"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

func (s *ClientService) GetPlayer(ctx context.Context, id int) (*playersapi.Player, error) {
	resp, err := s.client.GetPlayerById(ctx, playersapi.GetPlayerByIdParams{ID: id})
	if err != nil {
		log.Printf("[GET] Error: %v", err)
		return nil, err
	}
	if player, ok := resp.(*playersapi.Player); ok {
		log.Printf("[GET] Found player - ID=%d, Name=%s, Score=%.2f", player.ID, player.Name, player.Score)
		return player, nil
	}
	log.Printf("[GET] Unexpected response type: %T", resp)
	return nil, nil
}
