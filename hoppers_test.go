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

func TestMoveNotCorrectTurn(t *testing.T) {
	board := NewBoard()
	move := Move{Player: PlayerTwo, Piece: 16, Path: []Location{Location{7, 7}}}

	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should reject incorrect player turn")
	}

}

func TestMoveTooFar(t *testing.T) {
	board := NewBoard()
	move := Move{Player: PlayerOne, Piece: 26, Path: []Location{Location{7, 7}}}

	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should reject jump giant moves across the board")
	}

}

func TestMovePieceOutOfRange(t *testing.T) {
	board := NewBoard()
	move := Move{Player: PlayerOne, Piece: 31, Path: []Location{Location{7, 7}}}
	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should reject invalid Piece numbers")
	}

	board.CurrentTurn = PlayerTwo
	move = Move{Player: PlayerOne, Piece: 0, Path: []Location{Location{7, 7}}}
	err = board.Move(move)

	if err == nil {
		t.Errorf("Move() should reject invalid Piece numbers")
	}
}

func TestMovePieceGameWon(t *testing.T) {
	board := NewBoard()
	board.Winner = PlayerOne
	move := Move{Player: PlayerOne, Piece: 26, Path: []Location{Location{10, 5}}}
	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should reject moves after a player has won")
	}
}

func TestPlayerMovesOtherPlayerPiece(t *testing.T) {
	board := NewBoard()
	move := Move{Player: PlayerOne, Piece: 5, Path: []Location{Location{1, 6}}}
	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should reject attempts to move the other player's pieces")
	}

}

func TestValidMove(t *testing.T) {
	board := NewBoard()
	move := Move{Player: PlayerOne, Piece: 26, Path: []Location{Location{10, 5}}}
	err := board.Move(move)

	if err != nil {
		t.Errorf("Move() should accept valid moves", err)
	}

	if board.Pieces[26].Location != move.Path[0] {
		t.Errorf("Move() not updating piece positions correctly")
	}

	if board.Spaces[move.Path[0]] != move.Piece || board.Spaces[initialLocations[26]] != 0 {
		t.Errorf("Move() not updating space positions correctly")
	}
}

func TestPathSingleEntryAtStartLocation(t *testing.T) {
	board := NewBoard()
	// Piece 26 starts at Location{10, 6}
	move := Move{Player: PlayerOne, Piece: 26, Path: []Location{Location{10, 6}}}

	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should not allow a move to repeat locations")
	}
}

func TestEmptyPath(t *testing.T) {
	board := NewBoard()
	// Piece 26 starts at Location{10, 6}
	move := Move{Player: PlayerOne, Piece: 26, Path: []Location{}}

	err := board.Move(move)

	if err == nil {
		t.Errorf("Move() should not allow an empty Move.Path")
	}
}

func TestOutOfBoundsPaths(t *testing.T) {
	board := NewBoard()

	// XXX: this is meant to test invalid location checking
	// therefore that check must happen before hop logic checks
	paths := [][]Location{
		[]Location{Location{0, 1}},
		[]Location{Location{-1, 1}},
		[]Location{Location{11, 1}},
		[]Location{Location{1, 0}},
		[]Location{Location{1, -1}},
		[]Location{Location{1, 11}},
		[]Location{Location{0, 0}},
		[]Location{Location{-1, -1}},
		[]Location{Location{11, 11}},
	}

	for _, path := range paths {
		// Piece 26 starts at Location{10, 6}
		move := Move{Player: PlayerOne, Piece: 26, Path: path}
		err := board.Move(move)

		if err == nil {
			t.Errorf("Move() should not allow an empty Move.Path", path)
		}
	}
}
