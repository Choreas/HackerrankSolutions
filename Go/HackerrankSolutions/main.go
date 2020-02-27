package main

import (
	"fmt"

	"github.com/Choreas/HackerrankSolutions/HackerrankSolutions/Algorithms/implementation"
)

func main() {
	drawingBook()
}

func drawingBook() {
	fmt.Println("Drawing Book")

	var result int32
	var pages int32 = 5
	var pagesToTurn int32 = 4
	result = implementation.PageCount(pages, pagesToTurn)
	fmt.Println(result)
}
