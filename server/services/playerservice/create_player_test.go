package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (p *PlayerServiceSuite) TestCreatePlayerSuccess() {
	wantName := "Alice"
	wantScore := 5.5
	expectedID := 42

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.Name == wantName && player.Score == wantScore && player.ID == 0
	})).
		Return(&models.Player{ID: expectedID, Name: wantName, Score: wantScore}, nil).
		Once()

	gotPlayer, err := p.playerService.CreatePlayer(p.ctx, wantName, wantScore)

	assert.NoError(p.T(), err)
	assert.Equal(p.T(), expectedID, gotPlayer.ID)
	assert.Equal(p.T(), wantName, gotPlayer.Name)
	assert.Equal(p.T(), wantScore, gotPlayer.Score)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestCreatePlayerStorageError() {
	name := "Alice"
	score := 5.5
	wantErr := errors.New("storage error")

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.Name == name && player.Score == score && player.ID == 0
	})).
		Return(nil, wantErr).
		Once()

	gotPlayer, gotErr := p.playerService.CreatePlayer(p.ctx, name, score)

	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertExpectations(p.T())
}
