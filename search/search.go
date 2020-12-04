package search

import (
	"errors"
	"fmt"
)

type Interface interface {
	LinearSearch(search string, items []string) (string, error)
	Binary(search int, items []int) (int, error)
}

type search struct {
	Interface
}

func NewSearch() *search {
	return &search{}
}

func (s *search) LinearSearch(search string, items []string) (string, error) {
	for _, item := range items {
		if search == item {
			return item, nil
		}
	}

	return "", errors.New(fmt.Sprintf("no such item \"%s\" was found", search))
}

func (s *search) Binary(search int, items []int) (int, error) {
	min := 0
	max := len(items) - 1

	for min <= max {
		target := (min + max) / 2

		if items[target] == search {
			return items[target], nil
		}

		if items[target] < search {
			min = target + 1
		}

		if items[target] > search {
			max = target - 1
		}
	}

	return -1, errors.New(fmt.Sprintf("error: unable to locate \"%v\"", search))
}
