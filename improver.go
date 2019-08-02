package main

import "math/rand"

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
	worstBoard := i.Boards[0]
	worstLoss := i.Boards[0].Nearness()
	for _, b := range i.Boards {
		improveBoard(b)
		if b.Nearness() > worstLoss {
			worstBoard = b
			worstLoss = b.Nearness()
		}
	}
	if rand.Intn(10) == 0 {
		worstBoard.CopyFrom(i.BestBoard())
		for i := 0; i < rand.Intn(50); i++ {
			worstBoard.Mutate()
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
	for i := 0; i < 1000; i++ {
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
