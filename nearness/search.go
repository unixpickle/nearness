package nearness

import (
	"runtime"
)

const (
	LocalSteps = 20000
)

// Search looks for an improvement to the board b by
// running n parallel local searches and returning the
// best result.
//
// If n is 0, GOMAXPROCS is used.
func Search(b *Board, n int) *Board {
	if n == 0 {
		n = runtime.GOMAXPROCS(0)
	}

	boards := make(chan *Board, n)

	for i := 0; i < n; i++ {
		go func(first bool) {
			b1 := b.Copy()
			if !first {
				// TODO: test if this is even helpful.
				b1.Mutate()
			}
			LocalSearch(b1)
			boards <- b1
		}(i == 0)
	}

	bestBoard := b
	bestLoss := b.Nearness()
	for i := 0; i < n; i++ {
		b1 := <-boards
		if b1.Nearness() < bestLoss {
			bestBoard = b1
			bestLoss = b1.Nearness()
		}
	}

	return bestBoard
}

// LocalSearch searches for incremental improvements to
// the given board b.
//
// Returns true if the board was improved by the search.
func LocalSearch(b *Board) bool {
	improved := false
	b1 := b.Copy()
	for i := 0; i < LocalSteps; i++ {
		for j := 0; j < 5; j++ {
			b1.Mutate()
			if b1.Nearness() < b.Nearness() {
				improved = true
				b.CopyFrom(b1)
				break
			}
		}
		b1.CopyFrom(b)
	}
	return improved
}
