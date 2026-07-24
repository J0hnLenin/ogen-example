package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/config"
	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
	"github.com/J0hnLenin/ogen-example/server/internal/api/playersserviceapi"
	"github.com/rs/cors"
)

func AppRun(cfg config.ServerConfig, api *playersserviceapi.PlayerServiceAPI) {

	srv, err := playersapi.NewServer(api)
	if err != nil {
		log.Fatalf("Can't create server: %v", err)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	pattern := fmt.Sprintf("%s/", cfg.APIEndpoint)

	mux := http.NewServeMux()
	mux.Handle(pattern, http.StripPrefix(cfg.APIEndpoint, srv))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	log.Printf("Server listen: %s", addr)
	if err := http.ListenAndServe(addr, cors.Handler(mux)); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
