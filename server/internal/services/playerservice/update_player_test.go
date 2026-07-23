package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (p *PlayerServiceSuite) TestUpdatePlayerSuccess() {
	// Arrange
	id := 1
	newName := "NewAlice"
	newScore := 10.5
	oldPlayer := &models.Player{ID: id, Name: "Alice", Score: 5.5}
	updatedPlayer := &models.Player{ID: id, Name: newName, Score: newScore}

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(oldPlayer, nil).
		Once()

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.ID == id && player.Name == newName && player.Score == newScore
	})).
		Return(updatedPlayer, nil).
		Once()

	// Act
	gotPlayer, err := p.playerService.UpdatePlayer(p.ctx, id, newName, newScore)

	// Assert
	assert.NoError(p.T(), err)
	assert.Equal(p.T(), updatedPlayer, gotPlayer)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestUpdatePlayerGetError() {
	// Arrange
	id := 1
	newName := "NewAlice"
	newScore := 10.5
	wantErr := errors.New("player not found")

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(nil, wantErr).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.UpdatePlayer(p.ctx, id, newName, newScore)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertNotCalled(p.T(), "UpsertPlayer", mock.Anything, mock.Anything)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestUpdatePlayerUpsertError() {
	// Arrange
	id := 1
	newName := "NewAlice"
	newScore := 10.5
	oldPlayer := &models.Player{ID: id, Name: "Alice", Score: 5.5}
	wantErr := errors.New("storage error")

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(oldPlayer, nil).
		Once()

	p.storage.On("UpsertPlayer", p.ctx, mock.Anything).
		Return(nil, wantErr).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.UpdatePlayer(p.ctx, id, newName, newScore)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertExpectations(p.T())
}
