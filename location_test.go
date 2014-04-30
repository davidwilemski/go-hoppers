package hoppers

import "testing"

var LocationTests = []struct {
	L1, L2           Location // starting location and ending location
	RowDist, ColDist int      // Expected distance values
}{
	{Location{3, 3}, Location{2, 3}, 1, 0},
	{Location{3, 3}, Location{3, 2}, 0, 1},
	{Location{3, 3}, Location{2, 2}, 1, 1},
	{Location{3, 3}, Location{4, 4}, 1, 1},
	{Location{3, 3}, Location{1, 1}, 2, 2},
	{Location{3, 3}, Location{5, 5}, 2, 2},
}

func TestLocationDistance(t *testing.T) {
	for i, test := range LocationTests {
		if r, c := test.L1.Distance(test.L2); r != test.RowDist && c != test.ColDist {
			t.Errorf("Incorrect result on LocationTests[%d]: result: %d, %d", i, r, c)
		}
	}

}
