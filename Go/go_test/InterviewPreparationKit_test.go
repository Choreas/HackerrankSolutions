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

	for _, testcase := range cases {
		res := warmup.JumpingOnClouds(testcase.c)
		if res != testcase.exp {
			t.Errorf("%s: Result was incorrect, got: %d, expected: %d.", "JumpingOnClouds", res, testcase.exp)
		}
	}
}
