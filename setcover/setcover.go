package setcover

type SetCover interface {
	// FindPossibleSets is responsible for finding all possible sets that could match the U
	FindPossibleSets(subsets [][]int) [][]int
}

type setCover struct {
	SetCover
	universe []int
}

// NewSetCover constructs a new instance of SetCover.
func NewSetCover(universe []int) SetCover {
	return &setCover{
		universe: universe,
	}
}

func (sc *setCover) FindPossibleSets(subsets [][]int) [][]int {
	var s [][]int

	for _, ui := range sc.universe {
		for _, subset := range subsets {
			if sc.searchForMatch(subset, ui) {
				s = append(s, subset)
			}
		}
	}

	return s
}

// searchForMatch returns true iff there is an exact match with the subset items against the universe.
func (sc *setCover) searchForMatch(subset []int, needle int) bool {
	var matched bool

	for _, ssi := range subset {
		if ssi == needle {
			matched = true
			continue
		}
		matched = false
	}

	return matched
}
