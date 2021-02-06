package knapsack_test

import (
	"github.com/scrocrix/algorithms/knapsack"
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type knapsackUnitTest struct {
	suite.Suite
}

func (unit *knapsackUnitTest) TestNewKnapsack() {
	unit.Run("Return an instance of Knapsack interface", func() {
		sut := knapsack.NewKnapsack(10, new(insertionSortMock))
		require.NotNil(unit.T(), sut)
	})
}

func (unit *knapsackUnitTest) TestFirstFit() {
	unit.Run("Return items added from set to Knapsack", func() {
		sut := knapsack.NewKnapsack(10, new(insertionSortMock))
		result := sut.FirstFit([]int{5, 5})
		require.NotNil(unit.T(), result)
		require.NotEmpty(unit.T(), result)
		require.Equal(unit.T(), 2, len(sut.Items))
	})

	unit.Run("Return a set which is less than the limit", func() {
		sut := knapsack.NewKnapsack(10, new(insertionSortMock))
		result := sut.FirstFit([]int{5, 6})
		require.NotNil(unit.T(), result)
		require.NotEmpty(unit.T(), result)
		require.Equal(unit.T(), 1, len(sut.Items))
	})
}

func (unit *knapsackUnitTest) TestBestFit() {
	unit.Run("Return items which are ordered by its lowest value", func() {
		sortMock := new(insertionSortMock)

		sortMock.On("Sort", []int{5, 2, 3, 1}, sort.OrderAscending).Return([]int{1, 2, 3, 5}, nil)

		sut := knapsack.NewKnapsack(10, sortMock)

		result, err := sut.BestFit([]int{5, 2, 3, 1})

		require.Nil(unit.T(), err)
		require.NotNil(unit.T(), result)
		require.NotEmpty(unit.T(), result)

		require.Equal(unit.T(), []int{1, 2, 3}, sut.Items)
		require.Equal(unit.T(), result, sut.Items)
	})
}

func TestKnapsackUnitTest(t *testing.T) {
	suite.Run(t, new(knapsackUnitTest))
}

type knapsackIntegrationTest struct {
	suite.Suite
}

func (integration *knapsackIntegrationTest) TestBestFit() {
	integration.Run("Return items which are ordered by its lowest value", func() {
		sut := knapsack.NewKnapsack(10, sort.NewInsertion())

		result, err := sut.BestFit([]int{5, 2, 1, 3})

		require.Nil(integration.T(), err)
		require.NotEmpty(integration.T(), result)
		require.Equal(integration.T(), []int{1, 2, 3}, result)
		require.Equal(integration.T(), result, sut.Items)
	})
}

func TestKnapsackIntegrationTest(t *testing.T) {
	suite.Run(t, new(knapsackIntegrationTest))
}
