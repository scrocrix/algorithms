package knapsack_test

import (
	"github.com/scrocrix/algorithms/knapsack"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type knapsackUnitTest struct {
	suite.Suite
}

func (unit *knapsackUnitTest) TestNewKnapsack() {
	unit.Run("Return an instance of Knapsack interface", func() {
		sut := knapsack.NewKnapsack(10)
		require.NotNil(unit.T(), sut)
	})
}

func (unit *knapsackUnitTest) TestFirstFit() {
	unit.Run("Return items added from set to Knapsack", func() {
		sut := knapsack.NewKnapsack(10)
		result := sut.FirstFit([]int{5, 5})
		require.NotNil(unit.T(), result)
		require.NotEmpty(unit.T(), result)
		require.Equal(unit.T(), 2, len(sut.Items))
	})

	unit.Run("Return a set which is less than the limit", func() {
		sut := knapsack.NewKnapsack(10)
		result := sut.FirstFit([]int{5, 6})
		require.NotNil(unit.T(), result)
		require.NotEmpty(unit.T(), result)
		require.Equal(unit.T(), 1, len(sut.Items))
	})
}

func TestKnapsackUnitTest(t *testing.T) {
	suite.Run(t, new(knapsackUnitTest))
}
