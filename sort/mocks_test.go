package sort_test

import (
	"github.com/scrocrix/algorithms/sort"
	"github.com/stretchr/testify/mock"
)

type insertionMock struct {
	mock.Mock
	sort.Insertion
}

func (mock *insertionMock) Sort(items []int, orderBy string) ([]int, error) {
	args := mock.Called(items, orderBy)

	return args.Get(0).([]int), args.Error(1)
}
