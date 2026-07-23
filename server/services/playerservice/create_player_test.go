package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (p *PlayerServiceSuite) TestCreatePlayerSuccess() {
	// Arrange
	wantName := "Alice"
	wantScore := 5.5

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.Name == wantName && player.Score == wantScore && player.ID == 0
	})).
		Return(nil).
		Once()

	// Act
	gotPlayer, err := p.playerService.CreatePlayer(p.ctx, wantName, wantScore)

	// Assert
	assert.NoError(p.T(), err)
	assert.NotZero(p.T(), gotPlayer.ID, "ID should be generated")
	assert.Equal(p.T(), wantName, gotPlayer.Name)
	assert.Equal(p.T(), wantScore, gotPlayer.Score)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestCreatePlayerStorageError() {
	// Arrange
	name := "Alice"
	score := 5.5
	wantErr := errors.New("storage error")

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.Name == name && player.Score == score && player.ID == 0
	})).
		Return(wantErr).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.CreatePlayer(p.ctx, name, score)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertExpectations(p.T())
}
