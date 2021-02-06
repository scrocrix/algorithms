package knapsack_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/mock"
)

type insertionSortMock struct {
	mock.Mock
	sort.Insertion
}

func (mock *insertionSortMock) Sort(items []int, orderBy string) ([]int, error) {
	args := mock.Called(items, orderBy)

	return args.Get(0).([]int), args.Error(1)
}
