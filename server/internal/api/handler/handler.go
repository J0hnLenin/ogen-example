package handler

import "github.com/J0hnLenin/ogen-example/server/services/playerservice"

type PlayerHandler struct {
	service *playerservice.PlayerService
}

func NewPlayerHandler(service *playerservice.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}
