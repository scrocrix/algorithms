package sort

import (
	"errors"
)

var (
	OrderAscending  = "asc"
	OrderDescending = "desc"
	ErrEmptyParameter = "parameter must not be empty"
)

type Sort interface {
	// InsertionSort return an indefinite slice of ordered numbers. The order of these numbers might be
	// either Ascending or Descending, depending on the context.
	InsertionSort(items []int, orderBy string) ([]int, error)

	// MergeSort is an efficient way to sort a large slice by separating the slice into two slices,
	// individually sort them by making use of InsertionSort and then merging them together, therefore,
	// "Merge" and "Sort".
	MergeSort(item []int) []int
}

type sort struct {
	Sort
}

// NewSort constructs a new instance of Sort.
func NewSort() Sort {
	return &sort{}
}

func (s *sort) InsertionSort(items []int, orderBy string) ([]int, error) {
	if len(orderBy) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	for key, item := range items {
		o := key - 1

		if orderBy == OrderAscending {
			for o >= 0 && items[o] > item {
				items[o+1] = items[o]
				o = o - 1
			}
		}

		if orderBy == OrderDescending {
			for o >= 0 && items[o] < item {
				items[o+1] = items[o]
				o = o - 1
			}
		}

		items[o+1] = item
	}

	return items, nil
}

func (s *sort) MergeSort(items []int) []int {
	io := items[0:len(items)/2]
	it := items[len(items)/2:]

	ios, _ := s.InsertionSort(io, OrderAscending)
	its, _ := s.InsertionSort(it, OrderAscending)
	si, _ := s.InsertionSort(append(ios, its...), OrderAscending)

	return si
}
