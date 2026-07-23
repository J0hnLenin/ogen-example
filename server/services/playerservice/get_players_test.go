package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
	"github.com/stretchr/testify/assert"
)

func (p *PlayerServiceSuite) TestGetPlayersSuccess() {
	// Arrange
	wantPlayers := []*models.Player{
		{ID: 1, Name: "Alice", Score: 5.5},
		{ID: 2, Name: "Bob", Score: 8.0},
	}

	p.storage.On("GetPlayers", p.ctx).
		Return(wantPlayers, nil).
		Once()

	// Act
	gotPlayers, err := p.playerService.GetPlayers(p.ctx)

	// Assert
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), wantPlayers, gotPlayers)
}

func (p *PlayerServiceSuite) TestGetPlayersStorageError() {
	// Arrange
	wantErrorString := "storage error"
	wantErr := errors.New(wantErrorString)

	p.storage.On("GetPlayers", p.ctx).
		Return(nil, wantErr).
		Once()

	// Act
	gotPlayers, gotErr := p.playerService.GetPlayers(p.ctx)

	// Assert
	assert.Nil(p.T(), gotPlayers)
	assert.Equal(p.T(), wantErr, gotErr)
}

func (p *PlayerServiceSuite) TestGetPlayersEmpty() {
	// Arrange
	wantPlayers := []*models.Player{}

	p.storage.On("GetPlayers", p.ctx).
		Return(wantPlayers, nil).
		Once()

	// Act
	gotPlayers, err := p.playerService.GetPlayers(p.ctx)

	// Assert
	assert.Nil(p.T(), err)
	assert.Empty(p.T(), gotPlayers)
}
