package percolation_test

import (
	"github.com/scrocrix/algorithms/percolation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type percolationSuite struct {
	suite.Suite
}

func (suite *percolationSuite) TestNewPercolationReturnFiveByFiveSiteGrid() {
	result, _ := percolation.NewPercolation(5, 5)

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), 25, len(result.GetSites()))
}

func (suite *percolationSuite) TestNewPercolationReturnErrorWhenSiteSizeDoesNotMatch() {
	result, err := percolation.NewPercolation(5, 10)

	assert.Nil(suite.T(), result)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: row and column must be equal in order to create a quadratic percolation", err.Error())
}

func (suite *percolationSuite) TestOpenReturnTrueWhenSiteHasBeenOpened() {
	sut, _ := percolation.NewPercolation(5, 5)

	result, _ := sut.Open(1, 2)

	assert.True(suite.T(), result)
}

func (suite *percolationSuite) TestOpenReturnErrorWhenRowOrColumnAreLessThanZero() {
	sut, _ := percolation.NewPercolation(5, 5)

	isOpened, err := sut.Open(0, 1)

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: either row or column arguments must be greater than zero", err.Error())

	isOpened, err = sut.Open(1, 0)

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: either row or column arguments must be greater than zero", err.Error())
}

func (suite *percolationSuite) TestOpenReturnErrorWhenRowOrColumnAreGreaterThanPercolationSize() {
	sut, _ := percolation.NewPercolation(5, 5)

	isOpened, err := sut.Open(26, 2)

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: either row or column are greater than total sites count", err.Error())

	isOpened, err = sut.Open(3, 12)

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: either row or column are greater than total sites count", err.Error())
}

func TestPercolationSuite(t *testing.T) {
	suite.Run(t, new(percolationSuite))
}
