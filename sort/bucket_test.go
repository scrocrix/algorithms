package sort_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type bucketsortUnitTest struct {
	suite.Suite
}

func (unit *bucketsortUnitTest) TestNewBucketsort() {
	unit.Run("Return instance of Bucketsort when successful", func() {
		insertionMock := new(insertionMock)

		sut := sort.NewBucketsort(insertionMock)

		require.NotNil(unit.T(), sut)
	})
}

func (unit *bucketsortUnitTest) TestSort() {
	unit.Run("Return sorted slice when successful", func() {
		insertionMock := new(insertionMock)

		insertionMock.On("Sort", []int{9, 1, 8, 3}, sort.OrderAscending).Return([]int{1, 3, 8, 9}, nil)
		insertionMock.On("Sort", []int{12}, sort.OrderAscending).Return([]int{12}, nil)
		insertionMock.On("Sort", []int{20}, sort.OrderAscending).Return([]int{20}, nil)
		insertionMock.On("Sort", []int{43}, sort.OrderAscending).Return([]int{43}, nil)
		insertionMock.On("Sort", []int{55}, sort.OrderAscending).Return([]int{55}, nil)

		sut := sort.NewBucketsort(insertionMock)

		result := sut.Sort([]int{9, 12, 20, 55, 43, 1, 8, 3})

		require.Equal(unit.T(), []int{1, 3, 8, 9, 12, 20, 43, 55}, result)
	})
}

func TestBucketSortUnitTest(t *testing.T) {
	suite.Run(t, new(bucketsortUnitTest))
}

type bucketsortIntegrationTest struct {
	suite.Suite
}

func (integration *bucketsortIntegrationTest) TestSort() {
	integration.Run("Return a sorted slice of integers", func() {
		sut := sort.NewBucketsort(sort.NewInsertion())

		result := sut.Sort([]int{9, 12, 20, 55, 43, 1, 8, 3})

		require.Equal(integration.T(), []int{1, 3, 8, 9, 12, 20, 43, 55}, result)
	})
}

func TestBucketsortIntegrationTest(t *testing.T) {
	suite.Run(t, new(bucketsortIntegrationTest))
}
