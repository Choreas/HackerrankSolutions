package main

import (
	"fmt"

	implementation "github.com/Choreas/HackerrankSolutions/HackerrankSolutions/Practice/Algorithms/Implementation/go"
	warmup "github.com/Choreas/HackerrankSolutions/HackerrankSolutions/Practice/InterviewPreparationKit/Warm-upChallenges"
)

func main() {
	jumpingOnClouds()
}

func drawingBook() {
	fmt.Println("Drawing Book")

	var result int32
	var pages int32 = 5
	var pagesToTurn int32 = 4
	result = implementation.PageCount(pages, pagesToTurn)
	fmt.Println(result)
}

func countingValleys() {
	fmt.Println("Counting Valleys")
	var n int32 = 8
	var s string = "UDDDUDUU"

	fmt.Println(implementation.CountingValleys(n, s))
}

func jumpingOnClouds() {
	fmt.Println("Jumping On Clouds")
	cases := []struct {
		c   []int32
		exp int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
	}

	for _, testcase := range cases {
		res := warmup.JumpingOnClouds(testcase.c)
		fmt.Println(res)
	}
}
