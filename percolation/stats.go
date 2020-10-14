package percolation

import (
	"math/rand"
)

type stats struct {
	percolation PInterface
}

// NewStats create a new instance of stats struct
func NewStats(percolation PInterface) *stats {
	return &stats{
		percolation: percolation,
	}
}

// SearchForPercolation randomly select a site from the percolation grid
// it returns true when the process of searching for percolation has been completed.
//
// It doesn't necessarily mean that the percolation happened, only that the
// process of searching for percolation was completed though.
func (s *stats) SearchForPercolation() (bool, error) {
	sitesCount := len(s.percolation.GetSites())

	// select a random site from the percolation grid to start with
	randomPosition := rand.Intn(sitesCount-0) + 0

	i := 0

	for randomPosition < sitesCount && !s.percolation.GetSites()[randomPosition].IsOpened && i <= sitesCount {
		_, err := s.percolation.Open(s.percolation.GetSites()[randomPosition])

		if err != nil {
			return false, err
		}

		// select the next random site grid to open
		randomPosition = rand.Intn(sitesCount) + 0
		i++
	}

	return true, nil
}
