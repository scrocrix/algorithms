package insertionsort_test

import (
	"github.com/Scrocrix/algorithms/insertionsort"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type insertionSortSuite struct {
	suite.Suite
}

func (suite *insertionSortSuite) TestReturnSortedArrayInDescendingOrder() {
	sut := insertionsort.NewInsertionSort()

	un := []int{5, 4, 2, 6, 1, 3}
	result := sut.SortNumbers(un)

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), len(un), len(result))
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5, 6}, result)
}

func TestInsertionSortSuite(t *testing.T) {
	suite.Run(t, new(insertionSortSuite))
}