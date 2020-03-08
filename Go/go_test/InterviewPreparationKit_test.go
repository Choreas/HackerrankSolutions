package go_test

import (
	"testing"

	warmup "github.com/Choreas/HackerrankSolutions/HackerrankSolutions/Practice/InterviewPreparationKit/Warm-upChallenges"
)

func TestJumpingOnClouds(t *testing.T) {

	cases := []struct {
		c   []int32
		exp int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
	}

	for _, tc := range cases {
		res := warmup.JumpingOnClouds(tc.c)
		if res != tc.exp {
			t.Errorf("%s: Result was incorrect, got: %d, expected: %d.", "JumpingOnClouds", res, tc.exp)
		}
	}
}

func TestRepeatedString(t *testing.T) {

	cases := []struct {
		s   string
		n   int64
		exp int64
	}{
		{"aba", 10, 7},
		{"a", 1000000000000, 1000000000000},
		{"bkjh", 560, 0},
		{"", 10, 0},
	}

	for _, tc := range cases {
		res := warmup.RepeatedString(tc.s, tc.n)
		if res != tc.exp {
			t.Errorf("%s: Result was incorrect, got: %d, expected: %d.", "Repeated String", res, tc.exp)
		}
	}
}
