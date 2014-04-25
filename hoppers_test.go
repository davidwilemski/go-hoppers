package hoppers

import "testing"

func TestBoardInitFirstMove(t *testing.T) {

	board := NewBoard()

	if board.CurrentTurn != PlayerOne {
		t.Errorf("Board initialized with incorrect first turn: %s", board.CurrentTurn)
	}
}

func TestNewLocation(t *testing.T) {
	/*	var l Location*/
	var err error

	_, err = NewLocation(0, 0)
	if err == nil {
		t.Errorf("Locations should not allow values outside of (1, 30)")
	}

	_, err = NewLocation(1, 0)
	if err == nil {
		t.Errorf("Locations should not allow values outside of (1, 30)")
	}

	_, err = NewLocation(0, 1)
	if err == nil {
		t.Errorf("Locations should not allow values outside of (1, 30)")
	}
}
