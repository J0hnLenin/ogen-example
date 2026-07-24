package clientservice

import (
	"context"
	"log"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

func (s *ClientService) GetPlayers(ctx context.Context) ([]playersapi.Player, error) {
	players, err := s.client.GetPlayers(ctx)
	if err != nil {
		log.Printf("[GET] Error: %v", err)
		return nil, err
	}
	log.Printf("[GET] Found %d players", len(players))
	for _, p := range players {
		log.Printf("  - ID=%d, Name=%s, Score=%.2f", p.ID, p.Name, p.Score)
	}
	return players, nil
}
