package playerservice

import (
	"errors"

	"github.com/J0hnLenin/ogen-example/server/models"
	"github.com/stretchr/testify/assert"
)

func (p *PlayerServiceSuite) TestGetPlayerByIDSuccess() {
	// Arrange
	id := 1
	wantPlayer := &models.Player{
		ID:    id,
		Name:  "Alice",
		Score: 5.5,
	}

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(wantPlayer, nil).
		Once()

	// Act
	gotPlayer, err := p.playerService.GetPlayerByID(p.ctx, id)

	// Assert
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), wantPlayer, gotPlayer)
}

func (p *PlayerServiceSuite) TestGetPlayerByIDStorageError() {
	// Arrange
	wantErr := errors.New("storage error")

	id := 1

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(nil, wantErr).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.GetPlayerByID(p.ctx, id)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
}

func (p *PlayerServiceSuite) TestGetPlayerByIDNotFound() {
	// Arrange
	wantErr := errors.New("player not found")

	id := 1

	p.storage.On("GetPlayerByID", p.ctx, id).
		Return(nil, nil).
		Once()

	// Act
	gotPlayer, gotErr := p.playerService.GetPlayerByID(p.ctx, id)

	// Assert
	assert.Nil(p.T(), gotPlayer)
	assert.Equal(p.T(), wantErr, gotErr)
}
