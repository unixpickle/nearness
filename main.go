package main

import "fmt"

func main() {
	board := NewBoard(30).Shuffle()
	for {
		l1 := board.NormNearness()
		fmt.Printf("loss = %.10f\n", l1)
		b1 := board.Copy()
		b1.Mutate()
		l2 := b1.NormNearness()
		if l2 < l1 {
			board = b1
		}
	}
}
