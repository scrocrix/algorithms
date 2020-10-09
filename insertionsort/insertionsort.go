package insertionsort

import "errors"

var (
	SortOrderAscending  = "asc"
	SortOrderDescending = "desc"

	ErrEmptyParameter = "parameter must not be empty"
)

type InsertionSort interface {
	SortNumbers(numbers []int) ([]int, error)
}

type insertionSort struct {
	InsertionSort
}

// NewInsertionSort initializes the insertionSort struct for sorting
func NewInsertionSort() *insertionSort {
	return &insertionSort{}
}

// SortNumbers return an indefinite slice of ordered numbers. The order of these numbers might be
// either Ascending or Descending, depending on the context.
func (i *insertionSort) SortNumbers(numbers []int, orderBy string) ([]int, error) {
	if len(orderBy) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	for key, number := range numbers {
		numbers = i.sort(numbers, orderBy, key, number)
	}

	return numbers, nil
}

func (i *insertionSort) sort(numbers []int, order string, key int, number int) []int {
	o := key - 1

	if order == SortOrderAscending {
		for o >= 0 && numbers[o] > number {
			numbers[o+1] = numbers[o]
			o = o - 1
		}
	}

	if order == SortOrderDescending {
		for o >= 0 && numbers[o] < number {
			numbers[o+1] = numbers[o]
			o = o - 1
		}
	}

	numbers[o+1] = number

	return numbers
}
