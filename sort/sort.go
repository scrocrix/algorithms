package sort

import "errors"

var (
	OrderAscending  = "asc"
	OrderDescending = "desc"

	ErrEmptyParameter = "parameter must not be empty"
)

type Interface interface {
	InsertionSortIntegers(items interface{}, orderBy string) ([]int, error)
	InsertionSortAlphabetically(letters []string) []string
}

type sort struct {
	Interface
}

// NewSort initializes the sortIntegers struct for sorting
func NewSort() *sort {
	return &sort{}
}

// InsertionSortIntegers return an indefinite slice of ordered numbers. The order of these numbers might be
// either Ascending or Descending, depending on the context.
func (i *sort) InsertionSortIntegers(items []int, orderBy string) ([]int, error) {
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

// InsertionSortAlphabetically return an indefinite slice of ordered letters.
func (i *sort) InsertionSortAlphabetically(items []string) []string {
	for key, item := range items {
		o := key - 1

		for o >= 0 && items[o] > item {
			items[o+1] = items[o]
			o = o - 1
		}

		items[o+1] = item
	}

	return items
}
