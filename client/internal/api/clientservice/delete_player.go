package clientservice

import (
	"context"
	"log"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

func (s *ClientService) DeletePlayer(ctx context.Context, id int) error {
	params := playersapi.DeletePlayerParams{ID: id}
	resp, err := s.client.DeletePlayer(ctx, params)
	if err != nil {
		log.Printf("[DELETE] Error: %v", err)
		return err
	}

	switch resp.(type) {
	case *playersapi.DeletePlayerNoContent:
		log.Printf("[DELETE] Player %d deleted", id)
	default:
		log.Printf("[DELETE] Unexpected response type: %T", resp)
	}
	return nil
}
