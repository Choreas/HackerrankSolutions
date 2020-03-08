package go_test

import (
	"testing"

	implementation "github.com/Choreas/HackerrankSolutions/HackerrankSolutions/Practice/Algorithms/Implementation/go"
)

func TestPageCount(t *testing.T) {
	var (
		pages  int32 = 5
		toTurn int32 = 4
		exp    int32 = 0
	)

	if res := implementation.PageCount(pages, toTurn); res != exp {
		t.Errorf("%v != %v\n", res, exp)
	}
}

func TestCountingValleys(t *testing.T) {
	var (
		n   int32  = 8
		s   string = "UDDDUDUU"
		exp int32  = 1
	)

	if res := implementation.CountingValleys(n, s); res != exp {
		t.Errorf("%v != %v\n", res, exp)
	}
}
