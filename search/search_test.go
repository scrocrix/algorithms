package search_test

import (
	"github.com/scrocrix/algorithms/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type searchTest struct {
	suite.Suite
}

func (suite *searchTest) TestLinearSearchReturnCorrectCity() {
	sut := search.NewSearch()

	cities := []string{"Dubai", "Paris", "New York", "Amsterdam"}
	result, err := sut.LinearSearch("Amsterdam", cities)

	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "Amsterdam", result)
}

func (suite *searchTest) TestLinearSearchReturnErrorWhenItemWasNotFound() {
	sut := search.NewSearch()

	cities := []string{"Dubai", "Paris", "New York", "Amsterdam"}
	result, err := sut.LinearSearch("Sao Paulo", cities)

	assert.Empty(suite.T(), result)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "no such item \"Sao Paulo\" was found", err.Error())
}

func (suite *searchTest) TestBinaryReturnSearchNeedleFromSlice() {
	sut := search.NewSearch()

	result, _ := sut.Binary(67, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})

	assert.Equal(suite.T(), 67, result)

	result, _ = sut.Binary(3, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})

	assert.Equal(suite.T(), 3, result)
}

func (suite *searchTest) TestBinaryReturnErrorWhenSearchNeedleWasNotFound() {
	sut := search.NewSearch()

	result, err := sut.Binary(100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})

	assert.Equal(suite.T(), -1, result)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "error: unable to locate \"100\"", err.Error())
}

func TestSearchTest(t *testing.T) {
	suite.Run(t, new(searchTest))
}
