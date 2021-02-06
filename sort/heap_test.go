package sort_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type minHeapUnitTest struct {
	suite.Suite
}

func (unit *minHeapUnitTest) TestSort() {
	unit.Run("Return sorted heap structure", func() {
		sut, err := sort.NewMinHeap([]int{4, 3, 7, 1, 8, 5})

		require.Nil(unit.T(), err)

		require.Equal(unit.T(), []int{8, 7, 5, 4, 3, 1}, sut.Sort())
	})

	unit.Run("Returne error when heap items are empty", func() {
		sut, err := sort.NewMinHeap([]int{})

		require.NotNil(unit.T(), err)

		require.Nil(unit.T(), sut)
	})
}

func TestMinHeapUnitTest(t *testing.T) {
	suite.Run(t, new(minHeapUnitTest))
}
