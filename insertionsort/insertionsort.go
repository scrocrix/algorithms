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
	numbers []int
}

// NewInsertionSort initializes the insertionSort struct for sorting
func NewInsertionSort() *insertionSort {
	return &insertionSort{}
}

// SortNumbers return an indefinite slice of descending ordered numbers
func (i *insertionSort) SortNumbers(numbers []int, orderBy string) ([]int, error) {
	if len(orderBy) == 0 {
		return nil, errors.New(ErrEmptyParameter)
	}

	i.numbers = numbers

	for key, number := range numbers {
		o := i.sort(orderBy, key, number)
		numbers[o+1] = number
	}

	return numbers, nil
}

func (i *insertionSort) sort(order string, key int, number int) int {
	o := key - 1

	if order == SortOrderAscending {
		for o >= 0 && i.numbers[o] > number {
			i.numbers[o+1] = i.numbers[o]
			o = o - 1
		}
	}

	if order == SortOrderDescending {
		for o >= 0 && i.numbers[o] < number {
			i.numbers[o+1] = i.numbers[o]
			o = o - 1
		}
	}

	return o
}
