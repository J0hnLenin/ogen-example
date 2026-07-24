package main

import "github.com/J0hnLenin/ogen-example/server/internal/bootstrap"

func main() {
	storage := bootstrap.InitMemoryStorage()
	playerService := bootstrap.InitPlayerService(storage)
	playerAPI := bootstrap.InitPlayerAPI(playerService)

	bootstrap.AppRun(playerAPI)
}
