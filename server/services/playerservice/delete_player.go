package playerservice

import "context"

func (p *PlayerService) DeletePlayer(ctx context.Context, id int) error {
	return p.storage.DeletePlayer(ctx, id)
}
