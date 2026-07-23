package playerservice

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

//go:generate mockery --name=PlayerStorage
type PlayerStorage interface {
	GetPlayerByID(ctx context.Context, id int) (*models.Player, error)
	GetPlayers(ctx context.Context) ([]*models.Player, error)
	UpsertPlayer(ctx context.Context, player *models.Player) (*models.Player, error)
	DeletePlayer(ctx context.Context, id int) error
}

type PlayerService struct {
	storage PlayerStorage
}

func NewPlayerService(ctx context.Context, storage PlayerStorage) *PlayerService {
	return &PlayerService{
		storage: storage,
	}
}
