package handler

import (
	"context"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
)

type PlayerService interface {
	CreatePlayer(ctx context.Context, name string, score float64) (*models.Player, error)
	DeletePlayer(ctx context.Context, id int) error
	GetPlayerByID(ctx context.Context, id int) (*models.Player, error)
	GetPlayers(ctx context.Context) ([]*models.Player, error)
	UpdatePlayer(ctx context.Context, id int, newName string, newScore float64) (*models.Player, error)
	UpdatePlayerName(ctx context.Context, id int, newName string) (*models.Player, error)
	UpdatePlayerScore(ctx context.Context, id int, newScore float64) (*models.Player, error)
}

type PlayerHandler struct {
	service PlayerService
}

func NewPlayerHandler(ctx context.Context, service PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}
