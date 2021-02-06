package sort_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type insertionUnitTest struct {
	suite.Suite
}

func (unit *insertionUnitTest) TestSort() {
	unit.Run("Integers return slice of integers in ascending order", func() {
		sut := sort.NewInsertion()

		un := []int{5, 4, 2, 6, 1, 3}
		result, _ := sut.Sort(un, sort.OrderAscending)

		require.NotEmpty(unit.T(), result)
		require.Equal(unit.T(), len(un), len(result))
		require.Equal(unit.T(), []int{1, 2, 3, 4, 5, 6}, result)
	})

	unit.Run("Integers return slice of integers in descending order", func() {
		sut := sort.NewInsertion()

		un := []int{5, 4, 2, 6, 1, 3}
		result, _ := sut.Sort(un, sort.OrderDescending)

		require.NotEmpty(unit.T(), result)
		require.Equal(unit.T(), len(un), len(result))
		require.Equal(unit.T(), []int{6, 5, 4, 3, 2, 1}, result)
	})

	unit.Run("Integers return error when order by parameter is empty", func() {
		sut := sort.NewInsertion()

		un := []int{5, 4, 2, 6, 1, 3}
		_, err := sut.Sort(un, "")

		require.NotNil(unit.T(), err)
		require.Equal(unit.T(), sort.ErrEmptyParameter, err.Error())
	})
}

func TestInsertionUnitTest(t *testing.T) {
	suite.Run(t, new(insertionUnitTest))
}
