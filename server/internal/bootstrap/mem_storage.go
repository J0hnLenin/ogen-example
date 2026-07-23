package bootstrap

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/storage/memstorage"
)

func InitMemoryStorage() *memstorage.MemoryStorage {
	return memstorage.NewMemoryStorage(context.Background())
}
