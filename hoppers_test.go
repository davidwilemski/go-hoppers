package hoppers

import "testing"

func TestBoardInitFirstMove(t *testing.T) {

	board := NewBoard()

	if board.CurrentTurn != PlayerOne {
		t.Errorf("Board initialized with incorrect first turn: %s", board.CurrentTurn)
	}
}
