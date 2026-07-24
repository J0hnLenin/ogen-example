package main

import (
	"log"
	"os"

	"github.com/J0hnLenin/ogen-example/server/config"
	"github.com/J0hnLenin/ogen-example/server/internal/bootstrap"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		log.Fatalf("Load config error, %v", err)
	}

	storage := bootstrap.InitMemoryStorage()
	playerService := bootstrap.InitPlayerService(storage)
	playerAPI := bootstrap.InitPlayerAPI(playerService)

	bootstrap.AppRun(cfg.Server, playerAPI)
}
