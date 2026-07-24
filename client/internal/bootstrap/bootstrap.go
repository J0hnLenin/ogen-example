package bootstrap

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/J0hnLenin/ogen-example/client/internal/api/clientservice"
	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
	"github.com/J0hnLenin/ogen-example/client/internal/config"
	"github.com/J0hnLenin/ogen-example/client/internal/gen"
)

func Run(cfg config.ClientConfig) {
	ctx := context.Background()

	client, err := playersapi.NewClient(cfg.ServerURL)
	if err != nil {
		log.Fatalf("Failed to create API client: %v", err)
	}

	svc := clientservice.NewClientService(client)

	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()

	log.Printf("Client started. Server: %s, interval: %v", cfg.ServerURL, cfg.Interval)

	for {
		select {
		case <-ticker.C:
			doRandomOperation(ctx, svc)
		case <-ctx.Done():
			log.Println("Client stopped")
			return
		}
	}
}

func doRandomOperation(ctx context.Context, svc *clientservice.ClientService) {
	// Выбираем случайное действие
	actions := []func(context.Context, *clientservice.ClientService){
		doCreate,
		doUpdate,
		doPatch,
		doDelete,
		doGet,
	}
	action := actions[rand.Intn(len(actions))]
	action(ctx, svc)
}

func doCreate(ctx context.Context, svc *clientservice.ClientService) {
	input := gen.RandomPlayerInput()
	_ = svc.CreatePlayer(ctx, input)
}

func doUpdate(ctx context.Context, svc *clientservice.ClientService) {
	id := gen.RandomID()
	input := gen.RandomPlayerInput()
	_ = svc.UpdatePlayer(ctx, id, input)
}

func doPatch(ctx context.Context, svc *clientservice.ClientService) {
	id := gen.RandomID()
	patch := gen.RandomPlayerPartial()
	_ = svc.PatchPlayer(ctx, id, patch)
}

func doDelete(ctx context.Context, svc *clientservice.ClientService) {
	id := gen.RandomID()
	_ = svc.DeletePlayer(ctx, id)
}

func doGet(ctx context.Context, svc *clientservice.ClientService) {
	_, _ = svc.GetPlayers(ctx)
}
