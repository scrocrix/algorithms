package percolation

import "errors"

type percolation struct {
	sites [][]int
}

// NewPercolation creates a new instances of percolation as well as, configures
// the percolation grids being N-by-N size.
func NewPercolation(row int, col int) (*percolation, error) {
	percolation := &percolation{}

	if err := percolation.buildSites(row, col); err != nil {
		return nil, err
	}

	return percolation, nil
}

// buildSites essentially generates the grids used in the percolation process.
func (p *percolation) buildSites(row int, col int) error {
	if row != col {
		return errors.New("error: row and column must be equal in order to create a quadratic percolation")
	}

	var sites [][]int

	currentRow := 1

	currentColumn := 1

	for siteCount := 0; siteCount < row * col; siteCount++ {
		var site []int

		// verify if we must break to a new row
		if siteCount % row == 0 {
			currentRow = (siteCount / row) + 1

			// after breaking into a new row, the column must be reset to its initial position
			currentColumn = 1
		}

		site = append(site, currentColumn)
		site = append(site, currentRow)
		currentColumn++

		sites = append(sites, site)
	}

	p.sites = sites

	return nil
}

// GetSites return the sites that have been initialized with the percolation.
func (p *percolation) GetSites() [][]int {
	return p.sites
}

// Open finds the site in the grid based on the coordinates and verifies if
// the spot must or must not be opened.
func (p *percolation) Open(row int, col int) bool {
	return true
}
