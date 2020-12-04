package setcover_test

import (
	"github.com/scrocrix/algorithms/setcover"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type setCoverUnitSuite struct {
	suite.Suite
}

func (unit *setCoverUnitSuite) TestFindPossibleSets() {
	unit.Run("Return all possible subsets that might match the universe", func() {
		sut := setcover.NewSetCover([]int{1,2,3,4,5})

		result := sut.FindPossibleSets([][]int{
			{1,2,3},
			{2,4},
			{3,4},
			{4,5},
			{7,8,9},
		})

		require.NotEmpty(unit.T(), result)

		require.Equal(unit.T(), [][]int{
			{1,2,3},
			{2,4},
			{3,4},
			{4,5},
		}, result)
	})

	unit.Run("Return no subset matching", func() {
		sut := setcover.NewSetCover([]int{1,2,3,4,5})

		result := sut.FindPossibleSets([][]int{
			{4,5,6},
			{6,7,8},
			{8,9,10},
		})

		require.Empty(unit.T(), result)
	})
}

func TestSetCoverUnitSuite(t *testing.T) {
	suite.Run(t, new(setCoverUnitSuite))
}
