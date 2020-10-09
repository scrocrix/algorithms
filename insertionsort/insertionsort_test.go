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

func (suite *insertionSortSuite) TestSortNumbersReturnSliceOfIntegersInAscendingOrder() {
	sut := insertionsort.NewInsertionSort()

	un := []int{5, 4, 2, 6, 1, 3}
	result, _ := sut.SortNumbers(un, insertionsort.SortOrderAscending)

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), len(un), len(result))
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5, 6}, result)
}

func (suite *insertionSortSuite) TestSortNumbersReturnSliceOfIntegersInDescendingOrder() {
	sut := insertionsort.NewInsertionSort()

	un := []int{5, 4, 2, 6, 1, 3}
	result, _ := sut.SortNumbers(un, insertionsort.SortOrderDescending)

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), len(un), len(result))
	assert.Equal(suite.T(), []int{6, 5, 4, 3, 2, 1}, result)
}

func (suite *insertionSortSuite) TestSortNumbersReturnErrorWhenOrderByParameterIsEmpty() {
	sut := insertionsort.NewInsertionSort()

	un := []int{5, 4, 2, 6, 1, 3}
	_, err := sut.SortNumbers(un, "")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), insertionsort.ErrEmptyParameter, err.Error())
}

func TestInsertionSortSuite(t *testing.T) {
	suite.Run(t, new(insertionSortSuite))
}
