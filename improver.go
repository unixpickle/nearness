package main

import "math"
import "math/rand"

const (
	ImproveSteps = 200000
)

type Improver struct {
	Boards []*Board
}

func NewImprover(b *Board, numBoards int) *Improver {
	res := &Improver{Boards: make([]*Board, numBoards)}
	for i := range res.Boards {
		res.Boards[i] = b.Copy()
		if i != 0 {
			res.Boards[i].Mutate()
		}
	}
	return res
}

func (i *Improver) Step() {
	for _, b := range i.Boards {
		if !improveBoard(b) && b != i.BestBoard() {
			b.CopyFrom(i.BestBoard())
			num := int(math.Exp(rand.Float64() * math.Log(float64(b.Size*b.Size)/10)))
			for j := 0; j < num; j++ {
				b.Mutate()
			}
			break
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

func improveBoard(b *Board) bool {
	improved := false
	b1 := b.Copy()
	for i := 0; i < ImproveSteps; i++ {
		for j := 0; j < 1+rand.Intn(10); j++ {
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
