package main

import "github.com/unixpickle/essentials"

// SearchSwap searches over the space of all possible
// swaps to find the swap that improve the board as much
// as possible.
//
// Returns the result of all the swaps (that improve the
// board), or nil if no swap improves the nearness loss of
// the board.
func SearchSwap(b *Board) *Board {
	contributions := map[Position]int{}
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			p := Position{Row: i, Col: j}
			contributions[p] = computeContribution(b, p)
		}
	}

	swaps := &topN{N: 1000}
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
				v2 := b.At(p2.Row, p2.Col)
				*v1, *v2 = *v2, *v1
				delta += computeContribution(b, p1)
				delta += computeContribution(b, p2)
				*v1, *v2 = *v2, *v1
				swaps.Push(delta, p1, p2)
			}
		}
	}

	b1 := b.Copy()
	for i, p1 := range swaps.P1 {
		p2 := swaps.P2[i]
		delta := computeDelta(b1, p1, p2)
		if delta < 0 {
			b1.Swap(p1, p2)
		}
	}
	if b1.Nearness() < b.Nearness() {
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

func computeDelta(b *Board, p1, p2 Position) int {
	res := -(computeContribution(b, p1) + computeContribution(b, p2))
	b.Swap(p1, p2)
	res += computeContribution(b, p1) + computeContribution(b, p2)
	b.Swap(p1, p2)
	return res
}

type topN struct {
	N      int
	Deltas []int
	P1     []Position
	P2     []Position
}

func (t *topN) Push(delta int, p1, p2 Position) {
	if len(t.Deltas) < t.N {
		t.Deltas = append(t.Deltas, delta)
		t.P1 = append(t.P1, p1)
		t.P2 = append(t.P2, p2)
		return
	}
	if delta >= t.Deltas[t.N-1] {
		return
	}
	t.Deltas = append(t.Deltas, delta)
	t.P1 = append(t.P1, p1)
	t.P2 = append(t.P2, p2)
	essentials.VoodooSort(t.Deltas, func(i, j int) bool {
		return t.Deltas[i] < t.Deltas[j]
	}, t.P1, t.P2)
	t.Deltas = t.Deltas[:t.N]
	t.P1 = t.P1[:t.N]
	t.P2 = t.P2[:t.N]
}
