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
	}

	for _, p := range board.Pieces {
		board.Spaces[p.Location] = p.Num
	}

	return
}
