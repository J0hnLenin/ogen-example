package memstorage

import (
	"context"
	"errors"
)

func (m *MemoryStorage) DeletePlayer(ctx context.Context, id int) error {
	if m.playerMap[id] == nil {
		return errors.New("storage error")
	}

	m.playerMap[id] = nil
	return nil
}
