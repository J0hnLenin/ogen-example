package bootstrap

import (
	"log"
	"net/http"

	"github.com/J0hnLenin/ogen-example/server/internal/api/playersapi"
	"github.com/J0hnLenin/ogen-example/server/internal/api/playersserviceapi"
	"github.com/rs/cors"
)

func AppRun(api *playersserviceapi.PlayerServiceAPI) {

	srv, err := playersapi.NewServer(api)
	if err != nil {
		log.Fatalf("не удалось создать сервер: %v", err)
	}

	addr := "0.0.0.0:8000"
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", srv))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	log.Printf("сервер запущен на %s", addr)
	if err := http.ListenAndServe(addr, cors.Handler(mux)); err != nil {
		log.Fatalf("ошибка работы сервера: %v", err)
	}
}
