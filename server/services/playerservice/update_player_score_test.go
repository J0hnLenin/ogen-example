package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (p *PlayerServiceSuite) TestUpdatePlayerScoreSuccess() {
	// Arrange
	id := 1
	newScore := 10.5
	oldPlayer := &models.Player{ID: id, Name: "Alice", Score: 5.5}
	updatedPlayer := &models.Player{ID: id, Name: "Alice", Score: newScore}

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(oldPlayer, nil).
		Once()

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.ID == id && player.Name == oldPlayer.Name && player.Score == newScore
	})).
		Return(nil).
		Once()

	// Act
	gotPlayer, err := p.playerService.UpdatePlayerScore(p.ctx, id, newScore)

	// Assert
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), updatedPlayer, gotPlayer)
}

func (p *PlayerServiceSuite) TestUpdatePlayerScoreGetError() {
	// Arrange
	id := 1
	newScore := 10.5
	wantErr := errors.New("player not found")

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(nil, nil).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.UpdatePlayerScore(p.ctx, id, newScore)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertNotCalled(p.T(), "UpsertPlayer", mock.Anything, mock.Anything)
}

func (p *PlayerServiceSuite) TestUpdatePlayerScoreUpsertError() {
	// Arrange
	id := 1
	newScore := 10.5
	oldPlayer := &models.Player{ID: id, Name: "Alice", Score: 5.5}
	wantErr := errors.New("storage error")

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(oldPlayer, nil).
		Once()

	p.storage.On("UpsertPlayer", p.ctx, mock.Anything).
		Return(wantErr).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.UpdatePlayerScore(p.ctx, id, newScore)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
}
