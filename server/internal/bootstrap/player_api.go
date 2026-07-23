package bootstrap

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/api/handler"
	"github.com/J0hnLenin/ogen-example/server/internal/services/playerservice"
)

func InitPlayerAPI(playerService *playerservice.PlayerService) *handler.PlayerHandler {
	return handler.NewPlayerHandler(context.Background(), playerService)
}
