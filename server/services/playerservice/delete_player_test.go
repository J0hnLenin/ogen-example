package playerservice

import (
	"errors"

	"github.com/stretchr/testify/assert"
)

func (p *PlayerServiceSuite) TestDeletePlayerSuccess() {
	// Arrange
	id := 1

	p.storage.On("DeletePlayer", p.ctx, id).
		Return(nil).
		Once()

	// Act
	err := p.playerService.DeletePlayer(p.ctx, id)

	// Assert
	assert.NoError(p.T(), err)
	p.storage.AssertExpectations(p.T())
}

func (p *PlayerServiceSuite) TestDeletePlayerStorageError() {
	// Arrange
	id := 1
	wantErr := errors.New("storage error")

	p.storage.On("DeletePlayer", p.ctx, id).
		Return(wantErr).
		Once()

	// Act
	gotErr := p.playerService.DeletePlayer(p.ctx, id)

	// Assert
	assert.Equal(p.T(), wantErr, gotErr)
	p.storage.AssertExpectations(p.T())
}
