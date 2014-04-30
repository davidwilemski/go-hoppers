package hoppers

import "errors"

// Location contains coordinates for a game board
type Location struct {
	Row int
	Col int
}

// NewLocation instantiates a Location struct after performing bounds checks
func NewLocation(row int, col int) (l Location, err error) {
	if row <= 0 || row > 30 {
		err = errors.New("hoppers: row value is out of bounds")
	}

	if col <= 0 || col > 30 {
		err = errors.New("hoppers: col value is out of bounds")
	}

	l = Location{Row: row, Col: col}
	return
}

func (l Location) distanceRow(l2 Location) int {
	dist := l.Row - l2.Row
	if dist < 0 {
		dist *= -1
	}
	return dist
}

func (l Location) distanceCol(l2 Location) int {
	dist := l.Col - l2.Col
	if dist < 0 {
		dist *= -1
	}
	return dist
}

// Distance returns the row and column distances between two locations
// TODO test!
func (l Location) Distance(l2 Location) (row, col int) {
	row = l.distanceRow(l2)
	col = l.distanceCol(l2)
	return
}
