package hoppers

import "errors"

const (
	// PlayerOne string used to identify player 1
	PlayerOne = "one"

	// PlayerTwo is a string used to identify player 2
	PlayerTwo = "two"
)

var initialLocations = map[int32]Location{
	1: Location{Row: 1, Col: 1},
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
