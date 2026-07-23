package playerservice

import (
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func (p *PlayerServiceSuite) TestDeletePlayerNotFound() {
	// Arrange
	id := 0
	wantErr := errors.New("storage error")

	// Act
	gotErr := p.playerService.DeletePlayer(p.ctx, id)

	// Assert
	assert.Error(p.T(), gotErr)
	assert.Equal(p.T(), gotErr, wantErr)
	p.storage.AssertNotCalled(p.T(), "DeletePlayer", mock.Anything, mock.Anything)
}
