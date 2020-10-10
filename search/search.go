package search

import (
	"errors"
	"fmt"
)

type Interface interface {
	LinearSearch(search string, items []string) (string, error)
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