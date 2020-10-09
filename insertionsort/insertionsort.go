package insertionsort

type InsertionSort interface {
	// SortNumbers return an indefinite slice of descending ordered numbers
	SortNumbers(numbers []int) []int
}

type insertionSort struct {
	InsertionSort
}

// NewInsertionSort initializes the insertionSort struct for sorting
func NewInsertionSort() *insertionSort {
	return &insertionSort{}
}

func (i *insertionSort) SortNumbers(numbers []int) []int {
	for key, number := range numbers {
		i := key - 1

		for i >= 0 && numbers[i] > number {
			numbers[i + 1] = numbers[i]
			i = i - 1
		}

		numbers[i + 1] = number
	}

	return numbers
}
