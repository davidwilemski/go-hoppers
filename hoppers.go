package hoppers

import "errors"

const (
	// PlayerOne string used to identify player 1
	PlayerOne = "one"

	// PlayerTwo is a string used to identify player 2
	PlayerTwo = "two"
)

// Map of piece numbers to initial board locations
var initialLocations = map[int]Location{
	1:  Location{Row: 1, Col: 1},
	2:  Location{Row: 1, Col: 2},
	3:  Location{Row: 1, Col: 3},
	4:  Location{Row: 1, Col: 4},
	5:  Location{Row: 1, Col: 5},
	6:  Location{Row: 2, Col: 1},
	7:  Location{Row: 2, Col: 2},
	8:  Location{Row: 2, Col: 3},
	9:  Location{Row: 2, Col: 4},
	10: Location{Row: 3, Col: 1},
	11: Location{Row: 3, Col: 2},
	12: Location{Row: 3, Col: 3},
	13: Location{Row: 4, Col: 1},
	14: Location{Row: 4, Col: 2},
	15: Location{Row: 5, Col: 1},
	16: Location{Row: 6, Col: 10},
	17: Location{Row: 7, Col: 9},
	18: Location{Row: 7, Col: 10},
	19: Location{Row: 8, Col: 8},
	20: Location{Row: 8, Col: 9},
	21: Location{Row: 8, Col: 10},
	22: Location{Row: 9, Col: 7},
	23: Location{Row: 9, Col: 8},
	24: Location{Row: 9, Col: 9},
	25: Location{Row: 9, Col: 10},
	26: Location{Row: 10, Col: 6},
	27: Location{Row: 10, Col: 7},
	28: Location{Row: 10, Col: 8},
	29: Location{Row: 10, Col: 9},
	30: Location{Row: 10, Col: 10},
}

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
func (l Location) Distance(l2 Location) (row, col int) {
	row = l.distanceRow(l2)
	col = l.distanceCol(l2)
	return
}

// Piece contains state for a single piece that is on the game board
type Piece struct {
	Player   string
	Num      int
	Location Location
}

// Board contains complete state for a game of Hoppers
type Board struct {
	CurrentTurn string
	Pieces      map[int]Piece
	Spaces      map[Location]int
	Winner      string
}

func initPieces() map[int]Piece {
	pieces := make(map[int]Piece)
	for i, k := range initialLocations {

		player := PlayerOne
		if i < 16 {
			player = PlayerTwo
		}

		pieces[i] = Piece{Num: i, Location: k, Player: player}
	}
	return pieces
}

func initSpaces() (spaces map[Location]int) {
	spaces = make(map[Location]int)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			loc, _ := NewLocation(i, j)
			spaces[loc] = 0
		}
	}
	return
}

// NewBoard returns an initialized instance of the Board struct
func NewBoard() (board Board) {
	board = Board{
		CurrentTurn: PlayerOne, // default to player one
		Pieces:      initPieces(),
		Spaces:      initSpaces(), // create all nessecary tiles and zero them
		Winner:      "",
	}

	for _, p := range board.Pieces {
		board.Spaces[p.Location] = p.Num
	}

	return
}

// A Move struct contains any relevant state needed to validate and performa move on a piece
type Move struct {
	Player string
	Piece  int
	Path   []Location
}

// returns an error if the move is not legal according to hopper's rules
func (b Board) checkMoves(m Move) error {
	if len(m.Path) < 1 {
		return errors.New("path must contain at least one move")
	}

	for i, l := range m.Path {
		if l.Row > 10 || l.Col > 10 || l.Row < 1 || l.Col < 1 {
			return errors.New("path has a location that is not valid")
		}

		prev := b.Pieces[m.Piece].Location
		if i > 0 {
			prev = m.Path[i-1]
		}

		if l == prev {
			return errors.New("path may not repeat locations sequentially") // bad description
		}

		dist_row, dist_col := prev.Distance(l)
		if dist_row > 2 || dist_col > 2 {
			return errors.New("piece must move a single space or complete a hop")
		}

		if len(m.Path) > 1 && (dist_row < 2 && dist_col < 2) {
			return errors.New("A multiple location move must be all hops over a piece")
		}
	}

	return nil
}

// Move performs the move action given a valid Move struct
func (b *Board) Move(m Move) error {
	if m.Player != b.CurrentTurn {
		return errors.New("not player's turn")
	}

	if m.Piece < 1 || m.Piece > 30 {
		return errors.New("invalid piece number")
	}

	if b.Winner == PlayerOne || b.Winner == PlayerTwo {
		return errors.New("game over")
	}

	if b.Pieces[m.Piece].Player != m.Player {
		return errors.New("the player cannot move that piece")
	}

	if err := b.checkMoves(m); err != nil {
		return err
	}

	p := b.Pieces[m.Piece]
	newloc := m.Path[len(m.Path)-1]

	b.Spaces[p.Location] = 0   // invalidate old location
	b.Spaces[newloc] = m.Piece // set new location
	p.Location = newloc        // update location on Piece struct
	b.Pieces[m.Piece] = p      // put piece back into Pieces map

	return nil
}
