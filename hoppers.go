package hoppers

import "errors"

const (
	// PlayerOne string used to identify player 1
	PlayerOne = "one"

	// PlayerTwo is a string used to identify player 2
	PlayerTwo = "two"
)

// Map of piece numbers to initial board locations
var initialLocations = map[int32]Location{
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
	13: Location{Row: 4, Col: 4},
	14: Location{Row: 4, Col: 4},
	15: Location{Row: 5, Col: 5},
	16: Location{Row: 6, Col: 6},
	17: Location{Row: 7, Col: 7},
	18: Location{Row: 7, Col: 7},
	19: Location{Row: 8, Col: 8},
	20: Location{Row: 8, Col: 8},
	21: Location{Row: 8, Col: 8},
	22: Location{Row: 9, Col: 9},
	23: Location{Row: 9, Col: 9},
	24: Location{Row: 9, Col: 9},
	25: Location{Row: 9, Col: 9},
	26: Location{Row: 10, Col: 10},
	27: Location{Row: 10, Col: 10},
	28: Location{Row: 10, Col: 10},
	29: Location{Row: 10, Col: 10},
	30: Location{Row: 10, Col: 10},
}

// Location contains coordinates for a game board
type Location struct {
	Row int32
	Col int32
}

// NewLocation instantiates a Location struct after performing bounds checks
func NewLocation(row int32, col int32) (l Location, err error) {
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
	Num      int32
	Location Location
}

// Board contains complete state for a game of Hoppers
type Board struct {
	CurrentTurn string
	Pieces      map[int32]Piece
}

func initPieces() map[int32]Piece {
	pieces := make(map[int32]Piece)
	for i, k := range initialLocations {

		player := PlayerOne
		if i < 16 {
			player = PlayerTwo
		}

		pieces[i] = Piece{Num: i, Location: k, Player: player}
	}
	return pieces
}

// NewBoard returns an initialized instance of the Board struct
func NewBoard() Board {
	return Board{
		CurrentTurn: PlayerOne,
		Pieces:      initPieces(),
	}
}
