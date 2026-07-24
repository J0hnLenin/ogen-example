package main

import (
	"log"
	"os"

	"github.com/J0hnLenin/ogen-example/client/internal/bootstrap"
	"github.com/J0hnLenin/ogen-example/client/internal/config"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		log.Fatalf("Load config error, %v", err)
	}

	bootstrap.Run(cfg.Client)
}
