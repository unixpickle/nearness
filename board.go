package main

import (
	"math"

	"github.com/unixpickle/essentials"
)

type Position struct {
	Row int
	Col int
}

type Board struct {
	Size      int
	Positions []Position
}

func NewBoard(size int) *Board {
	b := &Board{
		Size:      size,
		Positions: make([]Position, size*size),
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			*b.At(i, j) = Position{Row: i, Col: j}
		}
	}
	return b
}

func (b *Board) At(i, j int) *Position {
	return &b.Positions[i*b.Size+j]
}

func (b *Board) Distance(p1, p2 Position) float64 {
	return math.Pow(b.coordDistance(p1.Col, p2.Col), 2) +
		math.Pow(b.coordDistance(p1.Row, p2.Row), 2)
}

func (b *Board) coordDistance(x1, x2 int) float64 {
	d1 := essentials.AbsInt(x1 - x2)
	return float64(essentials.MinInt(d1, b.Size-d1))
}

func (b *Board) Nearness() float64 {
	var res float64
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			a1 := *b.At(i, j)
			b1 := Position{Row: i, Col: j}
			for k := 0; k < b.Size; k++ {
				for l := 0; l < b.Size; l++ {
					a2 := *b.At(i, j)
					b2 := Position{Row: k, Col: l}
					res += b.Distance(a1, a2) * b.Distance(b1, b2)
				}
			}
		}
	}
	return res
}

func (b *Board) NormNearness() float64 {
	table := map[int]float64{
		1:  0,
		2:  10,
		3:  72,
		4:  816,
		5:  3800,
		6:  16902,
		7:  52528,
		8:  155840,
		9:  381672,
		10: 902550,
		11: 1883244,
		12: 3813912,
		13: 7103408,
		14: 12958148,
		15: 22225500,
		16: 37474816,
		17: 60291180,
		18: 95730984,
		19: 146469252,
		20: 221736200,
		21: 325763172,
		22: 474261920,
		23: 673706892,
		24: 949783680,
		25: 1311600000,
		26: 1799572164,
		27: 2425939956,
		28: 3252444776,
		29: 4294801980,
		30: 5643997650,
	}
	return b.Nearness() - table[b.Size]
}
