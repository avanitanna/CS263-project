package main

import "fmt"

type SquareArea interface {
	result() int
}

type Square struct {
	side int
}

func (square Square) result() int {
	return square.side * square.side
}

func getResult(squareArea SquareArea) int {
	return squareArea.result()
}

func main() {
	square := Square{side: 5}

	fmt.Printf("Square area: %d", getResult(square))
}
