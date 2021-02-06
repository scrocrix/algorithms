package sort

import "errors"

var (
	OrderAscending  = "asc"
	OrderDescending = "desc"
)

type Insertion interface {
	// Sort return an indefinite slice of ordered numbers. The order of these numbers might be
	// either Ascending or Descending, depending on the context.
	Sort(items []int, orderBy string) ([]int, error)
}

type insertion struct {
	Insertion
}

func NewInsertion() Insertion {
	return &insertion{}
}

func (s *insertion) Sort(items []int, orderBy string) ([]int, error) {
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
