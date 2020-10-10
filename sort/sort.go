package sort

import "errors"

var (
	OrderAscending  = "asc"
	OrderDescending = "desc"

	ErrEmptyParameter = "parameter must not be empty"
)

type Interface interface {
	InsertionSort(numbers []int, orderBy string) ([]int, error)
	SortAlphabetically(letters []string) []string
}

type sort struct {
	Interface
}

// NewSort initializes the sort struct for sorting
func NewSort() *sort {
	return &sort{}
}

// InsertionSort return an indefinite slice of ordered numbers. The order of these numbers might be
// either Ascending or Descending, depending on the context.
func (i *sort) InsertionSort(items interface{}, orderBy string) ([]int, error) {
	if len(orderBy) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	for key, item := range items.([]int) {
		items = i.sort(items.([]int), orderBy, key, item)
	}

	return items.([]int), nil
}

// SortAlphabetically return an indefinite slice of ordered letters.
func (i *sort) SortAlphabetically(letters []string) []string {
	for key, letter := range letters {
		o := key - 1

		for o >= 0 && letters[o] > letter {
			letters[o + 1] = letters[o]
			o = o - 1
		}

		// as default, we place the letter in the same place
		letters[o + 1] = letter
	}

	return letters
}

func (i *sort) sort(items []int, order string, key int, item int) []int {
	o := key - 1

	if order == OrderAscending {
		for o >= 0 && items[o] > item {
			items[o+1] = items[o]
			o = o - 1
		}
	}

	if order == OrderDescending {
		for o >= 0 && items[o] < item {
			items[o+1] = items[o]
			o = o - 1
		}
	}

	items[o+1] = item

	return items
}
