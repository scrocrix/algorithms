package percolation_test

import (
	"github.com/scrocrix/algorithms/percolation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

// percolationMock represents the percolation interface used when
// interacting with the percolation site grids.
type percolationMock struct {
	percolation.PInterface
	mock.Mock
}

func (pm *percolationMock) GetSites() []percolation.Site {
	args := pm.Called()
	return args.Get(0).([]percolation.Site)
}

func (pm *percolationMock) Open(site percolation.Site) (bool, error) {
	args := pm.Called(site)
	return args.Bool(0), args.Error(1)
}

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

func (suite *percolationSuite) TestOpenReturnTrueWhenSiteHasJustBeenOpened() {
	sut, _ := percolation.NewPercolation(5, 5)

	s := sut.GetSites()[2]

	result, _ := sut.Open(s)

	assert.True(suite.T(), result)
}

func (suite *percolationSuite) TestOpenReturnFalseWhenSiteHasAlreadyBeenOpened() {
	sut, _ := percolation.NewPercolation(5, 5)

	result, _ := sut.Open(sut.GetSites()[0])

	assert.True(suite.T(), result)

	result, _ = sut.Open(sut.GetSites()[0])

	assert.False(suite.T(), result)
}

func (suite *percolationSuite) TestOpenReturnFalseWhenSiteHasBeenBlocked() {
	sut, _ := percolation.NewPercolation(5, 5)

	result, _ := sut.Open(percolation.Site{
		Row:       1,
		Column:    2,
		IsOpened:  false,
		IsBlocked: true,
	})

	assert.False(suite.T(), result)
}

func (suite *percolationSuite) TestOpenReturnErrorWhenRowOrColumnAreLessThanZero() {
	sut, _ := percolation.NewPercolation(5, 5)

	isOpened, err := sut.Open(percolation.Site{
		Row:    0,
		Column: 1,
	})

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: either row or column arguments must be greater than zero", err.Error())

	isOpened, err = sut.Open(percolation.Site{
		Row:    1,
		Column: 0,
	})

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: either row or column arguments must be greater than zero", err.Error())
}

func (suite *percolationSuite) TestOpenReturnErrorWhenSiteWasNotFoundInPercolation() {
	sut, _ := percolation.NewPercolation(5, 5)

	isOpened, err := sut.Open(percolation.Site{
		Row:      1,
		Column:   1,
		Position: 5,
	})

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: the given site does not match any of the percolation's", err.Error())
}

func (suite *percolationSuite) TestOpenReturnErrorWhenSitePositionIsGreaterThanTheTotalPercolationCount() {
	sut, _ := percolation.NewPercolation(5, 5)

	isOpened, err := sut.Open(percolation.Site{
		Row:      1,
		Column:   1,
		Position: 30,
	})

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: the given site does not match any of the percolation's", err.Error())
}

func (suite *percolationSuite) TestOpenReturnErrorWhenSitePositionIsNegative() {
	sut, _ := percolation.NewPercolation(5, 5)

	isOpened, err := sut.Open(percolation.Site{
		Row:      1,
		Column:   1,
		Position: -1,
	})

	assert.False(suite.T(), isOpened)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error: site must not have a negative position", err.Error())
}

func TestPercolationSuite(t *testing.T) {
	suite.Run(t, new(percolationSuite))
}
