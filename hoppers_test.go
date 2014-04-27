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

	if len(pieces) != 30 {
		t.Errorf("Initial game state must have 30 pieces", pieces)
	}

	for i, p := range pieces {
		if p.Player == "" || p.Num == 0 {
			t.Errorf("Piece %d is not initialized", i, p)
		}

		if p.Num < 1 || p.Num > 30 {
			t.Errorf("There are only 30 pieces and numbered (1,30)", p)
		}

		if i < 16 && p.Player != PlayerTwo {
			t.Errorf("Piece numbers under 16 belong to player two", p)
		}

		if i > 15 && p.Player != PlayerOne {
			t.Errorf("Piece numbers above 15 belong to player one", p)
		}

		if p.Location != initialLocations[i] {
			t.Errorf("Piece %d has incorrect initial location", i, p)
		}
	}
}

func TestInitSpaces(t *testing.T) {
	board := NewBoard()
	spaces := board.Spaces

	if len(spaces) != 100 {
		t.Errorf("spaces does not have 100 locations")
	}

	for l, i := range spaces {

		if i < 0 || i > 30 {
			t.Errorf("invalid piece number on board", i)
		}

		// check that no expected initial locations are unoccupied
		if _, ok := initialLocations[i]; ok && spaces[l] == 0 {
			t.Errorf("initial location %! is unoccupied", l)
			return
		}

		if p, pExists := board.Pieces[i]; pExists && spaces[l] != p.Num {
			t.Errorf("Expected Piece %d at Location %!, got %d", i, l, spaces[l], initialLocations[i])
		}
	}
}
