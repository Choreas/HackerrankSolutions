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

func TestElectronicsShop(t *testing.T) {

	cases := []struct {
		budget    int32
		keyboards []int32
		drives    []int32
		exp       int32
	}{
		{10, []int32{3, 1}, []int32{5, 2, 8}, 9},
		{5, []int32{4}, []int32{5}, -1},
	}

	for _, tc := range cases {
		res := implementation.GetMoneySpent(tc.keyboards, tc.drives, tc.budget)
		if res != tc.exp {
			t.Errorf("%s: Result was incorrect, got: %d, expected: %d.", "Electronics Shop", res, tc.exp)
		}
	}
}
