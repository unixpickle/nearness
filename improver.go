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
			res.Boards[i].RandomSwap()
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
				b.RandomSwap()
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
	b1 := SearchSwap(b)
	if b1 != nil {
		b.CopyFrom(b1)
		return true
	}
	return false
}
