package main

import "math/rand"

const (
	MutationFrac = 0.03
	ImproveSteps = 1000
)

type Improver struct {
	Boards []*Board
}

func NewImprover(b *Board, numBoards int) *Improver {
	res := &Improver{Boards: make([]*Board, numBoards)}
	for i := range res.Boards {
		res.Boards[i] = b.Copy()
	}
	return res
}

func (i *Improver) Step() {
	for _, b := range i.Boards {
		improveBoard(b)
	}
	if rand.Float64() < MutationFrac {
		best := i.BestBoard()
		for {
			idx := rand.Intn(len(i.Boards))
			if i.Boards[idx] != best {
				i.Boards[idx].CopyFrom(best)
				for j := 0; j < rand.Intn(best.Size*best.Size/10); j++ {
					i.Boards[idx].Mutate()
				}
				break
			}
		}
	}
}

func (i *Improver) BestBoard() *Board {
	bestBoard := i.Boards[0]
	bestLoss := i.Boards[0].Nearness()
	for _, b := range i.Boards {
		if b.Nearness() < bestLoss {
			bestBoard = b
			bestLoss = b.Nearness()
		}
	}
	return bestBoard
}

func improveBoard(b *Board) {
	b1 := b.Copy()
	for i := 0; i < ImproveSteps; i++ {
		for i := 0; i < 1+rand.Intn(3); i++ {
			b1.Mutate()
		}
		if b1.Nearness() < b.Nearness() {
			b.CopyFrom(b1)
		} else {
			b1.CopyFrom(b)
		}
	}
}
