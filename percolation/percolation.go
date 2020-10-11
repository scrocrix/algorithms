package percolation

import (
	"errors"
	"math/rand"
)

// Site represents a single cell from the percolation's grid.
type Site struct {
	Row       int
	Column    int
	IsOpened  bool
	IsBlocked bool
	Position  int
}

// Validate check if the arguments in given to Site is correct
func (s *Site) Validate() error {
	if s.Row <= 0 || s.Column <= 0 {
		return errors.New("error: either row or column arguments must be greater than zero")
	}

	if s.Position < 0 {
		return errors.New("error: site must not have a negative position")
	}

	return nil
}

type percolation struct {
	rows    int
	columns int
	sites   []Site
}

// NewPercolation creates a new instances of percolation as well as, configures
// the percolation grids being N-by-N size.
func NewPercolation(row int, col int) (*percolation, error) {
	percolation := &percolation{
		rows:    row,
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

	var sites []Site

	currentRow := 1

	currentColumn := 1

	for siteCount := 0; siteCount < p.rows*p.columns; siteCount++ {
		var site Site

		// verify if we must break to a new row
		if siteCount%p.rows == 0 {
			currentRow = (siteCount / p.rows) + 1

			// after breaking into a new row, the column must be reset to its initial position
			currentColumn = 1
		}

		site.Row = currentRow
		site.Column = currentColumn
		site.IsOpened = false
		site.Position = siteCount

		// randomly assign a blocked site to the percolation
		site.IsBlocked = rand.Int()%2 == 0

		currentColumn++

		sites = append(sites, site)
	}

	p.sites = sites

	return nil
}

// GetSites return the sites that have been initialized with the percolation.
func (p *percolation) GetSites() []Site {
	return p.sites
}

// Open finds the site in the grid based on the coordinates and verifies if the spot must or must not be opened.
func (p *percolation) Open(site Site) (bool, error) {
	if err := site.Validate(); err != nil {
		return false, err
	}

	if site.Position > len(p.sites) || p.sites[site.Position] != site {
		return false, errors.New("error: the given site does not match any of the percolation's")
	}

	if site.IsOpened || site.IsBlocked {
		return false, nil
	}

	site.IsOpened = true
	p.sites[site.Position] = site

	return site.IsOpened, nil
}
