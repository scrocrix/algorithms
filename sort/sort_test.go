package sort_test

import (
	"github.com/Scrocrix/algorithms/sort"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type insertionSortSuite struct {
	suite.Suite
}

func (suite *insertionSortSuite) TestInsertionSortReturnSliceOfIntegersInAscendingOrder() {
	sut := sort.NewSort()

	un := []int{5, 4, 2, 6, 1, 3}
	result, _ := sut.InsertionSort(un, sort.OrderAscending)

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), len(un), len(result))
	assert.Equal(suite.T(), []int{1, 2, 3, 4, 5, 6}, result)
}

func (suite *insertionSortSuite) TestInsertionSortReturnSliceOfIntegersInDescendingOrder() {
	sut := sort.NewSort()

	un := []int{5, 4, 2, 6, 1, 3}
	result, _ := sut.InsertionSort(un, sort.OrderDescending)

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), len(un), len(result))
	assert.Equal(suite.T(), []int{6, 5, 4, 3, 2, 1}, result)
}

func (suite *insertionSortSuite) TestInsertionSortReturnErrorWhenOrderByParameterIsEmpty() {
	sut := sort.NewSort()

	un := []int{5, 4, 2, 6, 1, 3}
	_, err := sut.InsertionSort(un, "")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), sort.ErrEmptyParameter, err.Error())
}

func (suite *insertionSortSuite) TestSortAlphabeticallyReturnSliceOfStrings() {
	sut := sort.NewSort()

	result := sut.SortAlphabetically([]string{"b", "a", "d", "e", "c"})

	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), []string{"a", "b", "c", "d", "e"}, result)
}

func TestInsertionSortSuite(t *testing.T) {
	suite.Run(t, new(insertionSortSuite))
}
