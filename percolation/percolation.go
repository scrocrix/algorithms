package percolation

import "errors"

// Grid Example: 5 by 5
//
// [[1, 1], [1, 2], [1, 3], [1, 4], [1, 5]
//  [2, 1], [2, 2], [2, 3], [2, 4], [2, 5]
//  [3, 1], [3, 2], [3, 3], [3, 4], [3, 5]
//  [4, 1], [4, 2], [4, 3], [4, 4], [4, 5]
//  [5, 1], [5, 2], [5, 3], [5, 4], [5, 5]]

type percolation struct {
	rows int
	columns int
	sites [][]int
}

// NewPercolation creates a new instances of percolation as well as, configures
// the percolation grids being N-by-N size.
func NewPercolation(row int, col int) (*percolation, error) {
	percolation := &percolation{
		rows: row,
		columns: col,
	}

	if err := percolation.buildSites(); err != nil {
		return nil, err
	}

	return percolation, nil
}

// buildSites essentially generates the grids used in the percolation process.
func (p *percolation) buildSites() error {
	if p.rows != p.columns {
		return errors.New("error: row and column must be equal in order to create a quadratic percolation")
	}

	var sites [][]int

	currentRow := 1

	currentColumn := 1

	for siteCount := 0; siteCount < p.rows * p.columns; siteCount++ {
		var site []int

		// verify if we must break to a new row
		if siteCount % p.rows == 0 {
			currentRow = (siteCount / p.rows) + 1

			// after breaking into a new row, the column must be reset to its initial position
			currentColumn = 1
		}

		site = append(site, currentRow)
		site = append(site, currentColumn)
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
func (p *percolation) Open(row int, col int) (bool, error) {
	if row <= 0 || col <= 0 {
		return false, errors.New("error: either row or column arguments must be greater than zero")
	}

	if row > p.rows || col > p.columns {
		return false, errors.New("error: either row or column are greater than total sites count")
	}

	return true, nil
}
