package implementation

import "testing"

func drawingBook(t *testing.T) {
	var (
		pages  int32 = 5
		toTurn int32 = 4
		exp    int32 = 0
	)

	if res := PageCount(pages, toTurn); res != exp {
		t.Errorf("%v != %v\n", res, exp)
	}
}
