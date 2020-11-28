package knapsack

// Knapsack aims to solve the "Knapsack Problem" by utilizing three fundamental combinatorial algorithms.
type Knapsack interface {
	// FirstFit is based on the concept that given a set "S", it adds new items into the knapsack until all items
	// out of the set are added or when the knapsack limit is reached.
	FirstFit(s []int) []int
}

type knapsack struct {
	Knapsack
	Limit int
	Items []int
}

// NewKnapsack constructs a new instance of Knapsack.
func NewKnapsack(limit int) *knapsack {
	return &knapsack{
		Limit: limit,
	}
}

func (k *knapsack) FirstFit(s []int) []int {
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


