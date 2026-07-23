package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (p *PlayerServiceSuite) TestUpdatePlayerNameSuccess() {
	id := 1
	newName := "NewAlice"
	oldPlayer := &models.Player{ID: id, Name: "Alice", Score: 5.5}
	updatedPlayer := &models.Player{ID: id, Name: newName, Score: 5.5}

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(oldPlayer, nil).
		Once()

	p.storage.On("UpsertPlayer", p.ctx, mock.MatchedBy(func(player *models.Player) bool {
		return player.ID == id && player.Name == newName && player.Score == oldPlayer.Score
	})).
		Return(updatedPlayer, nil).
		Once()

	gotPlayer, err := p.playerService.UpdatePlayerName(p.ctx, id, newName)

	assert.NoError(p.T(), err)
	assert.Equal(p.T(), updatedPlayer, gotPlayer)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestUpdatePlayerNameGetError() {
	id := 1
	newName := "NewAlice"
	wantErr := errors.New("player not found")

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(nil, wantErr).
		Once()

	gotPlayer, gotErr := p.playerService.UpdatePlayerName(p.ctx, id, newName)

	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertNotCalled(p.T(), "UpsertPlayer", mock.Anything, mock.Anything)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestUpdatePlayerNameUpsertError() {
	id := 1
	newName := "NewAlice"
	oldPlayer := &models.Player{ID: id, Name: "Alice", Score: 5.5}
	wantErr := errors.New("storage error")

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(oldPlayer, nil).
		Once()

	p.storage.On("UpsertPlayer", p.ctx, mock.Anything).
		Return(nil, wantErr).
		Once()

	gotPlayer, gotErr := p.playerService.UpdatePlayerName(p.ctx, id, newName)

	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertExpectations(p.T())
}
