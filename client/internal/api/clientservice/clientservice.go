package clientservice

import (
	"github.com/J0hnLenin/ogen-example/client/internal/api/playersapi"
)

type ClientService struct {
	client *playersapi.Client
}

func NewClientService(client *playersapi.Client) *ClientService {
	return &ClientService{client: client}
}
