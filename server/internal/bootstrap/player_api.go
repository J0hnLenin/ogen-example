package bootstrap

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersserviceapi"
	"github.com/J0hnLenin/ogen-example/server/internal/services/playerservice"
)

func InitPlayerAPI(playerService *playerservice.PlayerService) *playersserviceapi.PlayerServiceAPI {
	return playersserviceapi.NewPlayerServiceAPI(context.Background(), playerService)
}
