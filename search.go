package main

// SearchSwap searches over the space of all possible
// swaps to find the swap that improves the board as well
// as possible.
//
// Returns the result of the swap, or nil if no swap
// improves the nearness loss of the board.
func SearchSwap(b *Board) *Board {
	contributions := map[Position]int{}
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			p := Position{Row: i, Col: j}
			contributions[p] = computeContribution(b, p)
		}
	}

	var bestDelta int
	var bestPos1 Position
	var bestPos2 Position
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			p1 := Position{Row: i, Col: j}
			v1 := b.At(i, j)
			contrib1 := contributions[p1]
			for p2, contrib2 := range contributions {
				if !p1.LessThan(p2) {
					continue
				}
				delta := -(contrib1 + contrib2)
				if delta > bestDelta {
					// It's impossible for this to be a bigger
					// improvement at this point.
					continue
				}
				v2 := b.At(p2.Row, p2.Col)
				*v1, *v2 = *v2, *v1
				delta += computeContribution(b, p1)
				delta += computeContribution(b, p2)
				*v1, *v2 = *v2, *v1
				if delta < bestDelta {
					bestDelta = delta
					bestPos1 = p1
					bestPos2 = p2
				}
			}
		}
	}

	if bestDelta < 0 {
		b1 := b.Copy()
		v1 := b.At(bestPos1.Row, bestPos1.Col)
		v2 := b.At(bestPos2.Row, bestPos2.Col)
		*v1, *v2 = *v2, *v1
		return b1
	}

	return nil
}

func computeContribution(b *Board, p Position) int {
	v := *b.At(p.Row, p.Col)
	total := 0
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			p1 := Position{Row: i, Col: j}
			v1 := *b.At(i, j)
			total += b.Distance(p, p1) * b.Distance(v, v1)
		}
	}
	return total
}
