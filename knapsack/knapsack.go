package knapsack

import "github.com/scrocrix/algorithms/sort"

// Knapsack aims to solve the "Knapsack Problem" by utilizing three fundamental combinatorial algorithms.
type Knapsack interface {
	// FirstFit is based on the concept that given a set "S", it adds new items into the knapsack until all items
	// out of the set are added or when the knapsack limit is reached.
	FirstFit(s []int) []int

	// BestFit  goes one step further than FirstFit where it also optimizes the items which get added to
	// knapsack, where in this context, the set gets sorted by the lowest value.
	BestFit(s []int) []int
}

type knapsack struct {
	Knapsack
	Limit int
	Items []int
	sort  sort.Insertion
}

// NewKnapsack constructs a new instance of Knapsack.
func NewKnapsack(limit int, sort sort.Insertion) *knapsack {
	return &knapsack{
		Limit: limit,
		sort:  sort,
	}
}

// addItems receives a set of items as argument and adds it to the knapsack
// until it reaches its limit.
func (k *knapsack) addItems(s []int) []int {
	var addedToKnapsack int

	for _, si := range s {
		addedToKnapsack = addedToKnapsack + si

		if addedToKnapsack > k.Limit {
			return k.Items
		}

		k.Items = append(k.Items, si)
	}

	return k.Items
}

func (k *knapsack) FirstFit(s []int) []int {
	return k.addItems(s)
}

func (k *knapsack) BestFit(s []int) ([]int, error) {
	set, err := k.sort.Sort(s, sort.OrderAscending)

	if err != nil {
		return nil, err
	}

	return k.addItems(set), nil
}
