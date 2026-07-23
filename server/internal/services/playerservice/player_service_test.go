package playerservice

import (
	"context"
	"testing"

	"github.com/J0hnLenin/ogen-example/server/internal/services/playerservice/mocks"
	"github.com/stretchr/testify/suite"
)

type PlayerServiceSuite struct {
	suite.Suite
	ctx           context.Context
	storage       *mocks.PlayerStorage
	playerService *PlayerService
}

func (p *PlayerServiceSuite) SetupTest() {
	p.ctx = context.Background()
	p.storage = mocks.NewPlayerStorage(p.T())

	p.playerService = NewPlayerService(p.ctx, p.storage)
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerServiceSuite))
}
