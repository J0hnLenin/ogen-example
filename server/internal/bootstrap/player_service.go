package bootstrap

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/services/playerservice"
	"github.com/J0hnLenin/ogen-example/server/internal/storage/memstorage"
)

func InitPlayerService(storage *memstorage.MemoryStorage) *playerservice.PlayerService {
	return playerservice.NewPlayerService(context.Background(), storage)
}
