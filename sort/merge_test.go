package sort_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mergeUnitTest struct {
	suite.Suite
}

func (unit *mergeUnitTest) TestSort() {
	unit.Run("Return sorted items", func() {
		sut := sort.NewMerge()

		require.Equal(unit.T(), []int{4, 5, 8, 10}, sut.Sort([]int{8, 4, 5, 10}))
	})
}

func TestMergeUnitTest(t *testing.T) {
	suite.Run(t, new(mergeUnitTest))
}
