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

func TestSearchTest(t *testing.T) {
	suite.Run(t, new(searchTest))
}