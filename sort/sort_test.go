package sort_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type sortSuite struct {
	suite.Suite
}

func (suite *sortSuite) TestInsertionSort() {
	suite.Run("Integers return slice of integers in ascending order", func() {
		sut := sort.NewSort()

		un := []int{5, 4, 2, 6, 1, 3}
		result, _ := sut.InsertionSort(un, sort.OrderAscending)

		assert.NotEmpty(suite.T(), result)
		assert.Equal(suite.T(), len(un), len(result))
		assert.Equal(suite.T(), []int{1, 2, 3, 4, 5, 6}, result)
	})

	suite.Run("Integers return slice of integers in descending order", func() {
		sut := sort.NewSort()

		un := []int{5, 4, 2, 6, 1, 3}
		result, _ := sut.InsertionSort(un, sort.OrderDescending)

		assert.NotEmpty(suite.T(), result)
		assert.Equal(suite.T(), len(un), len(result))
		assert.Equal(suite.T(), []int{6, 5, 4, 3, 2, 1}, result)
	})

	suite.Run("Integers return error when order by parameter is empty", func() {
		sut := sort.NewSort()

		un := []int{5, 4, 2, 6, 1, 3}
		_, err := sut.InsertionSort(un, "")

		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), sort.ErrEmptyParameter, err.Error())
	})
}

func (suite *sortSuite) TestMergeSort() {
	suite.Run("Return sorted slice of integers when successful", func() {
		sut := sort.NewSort()

		result := sut.MergeSort([]int{17, 11, 5, 7, 3, 13, 2, 19})

		assert.Equal(suite.T(), []int{2, 3, 5, 7, 11, 13, 17, 19}, result)
	})
}

func TestSortSuite(t *testing.T) {
	suite.Run(t, new(sortSuite))
}
