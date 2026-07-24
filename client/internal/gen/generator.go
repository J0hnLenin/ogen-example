package gen

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomPlayerInput() *playersapi.PlayerInput {
	return &playersapi.PlayerInput{
		Name:  fmt.Sprintf("Player-%d", rng.Intn(10)),
		Score: rng.Float64() * 100,
	}
}

func RandomPlayerPartial() *playersapi.PlayerPartial {
	name := fmt.Sprintf("Player-%d", rng.Intn(10))
	score := rng.Float64() * 100
	if rng.Intn(2) == 0 {
		return &playersapi.PlayerPartial{
			Name:  playersapi.OptString{Value: name, Set: true},
			Score: playersapi.OptFloat64{Value: 0.0, Set: false},
		}
	} else {
		return &playersapi.PlayerPartial{
			Name:  playersapi.OptString{Value: "", Set: false},
			Score: playersapi.OptFloat64{Value: score, Set: true},
		}
	}
}

func RandomID() int {
	return rng.Intn(10) + 1
}
