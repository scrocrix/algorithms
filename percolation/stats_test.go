package percolation_test

import (
	"github.com/scrocrix/algorithms/percolation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type statsSuite struct {
	suite.Suite
}

func (suite *statsSuite) TestNewStatsToReturnANewInstanceOfStatsStruct() {
	p := new(percolationMock)

	sut := percolation.NewStats(p)

	assert.NotEmpty(suite.T(), sut)
}

func (suite *statsSuite) TestSearchForPercolationReturnTrueWhenSuccessful() {
	p := new(percolationMock)

	p.On("GetSites").Return([]percolation.Site{
		{Row: 1, Column: 1, IsOpened: false, IsBlocked: false, Position: 0},
		{Row: 1, Column: 2, IsOpened: false, IsBlocked: false, Position: 1},
		{Row: 1, Column: 3, IsOpened: false, IsBlocked: false, Position: 2},
		{Row: 2, Column: 1, IsOpened: false, IsBlocked: false, Position: 3},
		{Row: 2, Column: 2, IsOpened: false, IsBlocked: false, Position: 4},
		{Row: 2, Column: 3, IsOpened: false, IsBlocked: false, Position: 5},
		{Row: 3, Column: 1, IsOpened: false, IsBlocked: false, Position: 6},
		{Row: 3, Column: 2, IsOpened: false, IsBlocked: false, Position: 7},
		{Row: 3, Column: 3, IsOpened: false, IsBlocked: false, Position: 8},
	})

	p.On("Open", percolation.Site{Row: 1, Column: 1, IsOpened: false, IsBlocked: false, Position: 0}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 1, Column: 2, IsOpened: false, IsBlocked: false, Position: 1}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 1, Column: 3, IsOpened: false, IsBlocked: false, Position: 2}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 2, Column: 1, IsOpened: false, IsBlocked: false, Position: 3}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 2, Column: 2, IsOpened: false, IsBlocked: false, Position: 4}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 2, Column: 3, IsOpened: false, IsBlocked: false, Position: 5}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 3, Column: 1, IsOpened: false, IsBlocked: false, Position: 6}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 3, Column: 2, IsOpened: false, IsBlocked: false, Position: 7}).Return(true, nil)
	p.On("Open", percolation.Site{Row: 3, Column: 3, IsOpened: false, IsBlocked: false, Position: 8}).Return(true, nil)

	sut := percolation.NewStats(p)

	finishedPercolation, err := sut.SearchForPercolation()

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), finishedPercolation)
}

func TestStatsSuite(t *testing.T) {
	suite.Run(t, new(statsSuite))
}
