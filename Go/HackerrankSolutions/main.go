package main

import (
	"fmt"

	implementation "github.com/Choreas/HackerrankSolutions/HackerrankSolutions/Practice/Algorithms/Implementation/go"
)

func main() {
	countingValleys()
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
