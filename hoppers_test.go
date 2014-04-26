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

func TestInitPieces(t *testing.T) {
	pieces := initPieces()

	if len(pieces) == 0 {
		t.Errorf("No Pieces in game state", pieces)
	}

	if len(pieces) > 30 {
		t.Errorf("Too many Pieces in game state", pieces)
	}

	for i, p := range pieces {
		if p.Player == "" || p.Num == 0 {
			t.Errorf("Piece %d is not initialized", i, p)
		}

		if p.Num < 1 || p.Num > 30 {
			t.Errorf("There are only 30 pieces and numbered (1,30)", p)
		}

		if i < 16 && p.Player != PlayerTwo {
			t.Errorf("Piece numbers under 16 belong to player two")
		}

		if i > 15 && p.Player != PlayerOne {
			t.Errorf("Piece numbers above 15 belong to player one")
		}
	}
}
