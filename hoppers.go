package hoppers

const (
	// PlayerOne string used to identify player 1
	PlayerOne = "one"

	// PlayerTwo is a string used to identify player 2
	PlayerTwo = "two"
)

// Board contains complete state for a game of Hoppers
type Board struct {
	CurrentTurn string
}

// NewBoard returns an initialized instance of the Board struct
func NewBoard() Board {
	return Board{
		CurrentTurn: PlayerOne,
	}
}
