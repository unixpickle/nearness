package main

import (
	"math/rand"
	"strings"

	"github.com/unixpickle/essentials"
)

type Position struct {
	Row int
	Col int
}

func (p Position) LessThan(p1 Position) bool {
	return p.Row < p1.Row || (p.Row == p1.Row && p.Col < p1.Col)
}

func (p Position) String() string {
	return idxString(p.Col) + idxString(p.Row)
}

func idxString(idx int) string {
	if idx < 26 {
		return string('A' + rune(idx))
	} else {
		return string('1' + rune(idx-26))
	}
}

type Board struct {
	Size          int
	Positions     []Position
	nearnessCache int
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

func (b *Board) Copy() *Board {
	res := &Board{
		Size:      b.Size,
		Positions: make([]Position, b.Size*b.Size),

		nearnessCache: b.nearnessCache,
	}
	copy(res.Positions, b.Positions)
	return res
}

func (b *Board) CopyFrom(b1 *Board) {
	b.Size = b1.Size
	b.nearnessCache = b1.nearnessCache
	copy(b.Positions, b1.Positions)
}

func (b *Board) Mutate() {
	p1 := b.randomPos()
	p2 := b.randomPos()
	v1 := b.At(p1.Row, p1.Col)
	v2 := b.At(p2.Row, p2.Col)

	updateCache := func(sign int) {
		if b.nearnessCache == 0 {
			return
		}
		vs := []Position{*v1, *v2}
		for pointIdx, p := range []Position{p1, p2} {
			v := vs[pointIdx]
			for i := 0; i < b.Size; i++ {
				for j := 0; j < b.Size; j++ {
					p1 := Position{Row: i, Col: j}
					v1 := *b.At(i, j)
					b.nearnessCache += sign * b.Distance(p, p1) * b.Distance(v, v1)
				}
			}
		}
	}

	updateCache(-1)
	*v1, *v2 = *v2, *v1
	updateCache(1)
}

func (b *Board) randomPos() Position {
	return Position{
		Row: rand.Intn(b.Size),
		Col: rand.Intn(b.Size),
	}
}

func (b *Board) Shuffle() *Board {
	perm := rand.Perm(len(b.Positions))
	res := &Board{
		Size:      b.Size,
		Positions: make([]Position, b.Size*b.Size),
	}
	for i, j := range perm {
		res.Positions[i] = b.Positions[j]
	}
	return res
}

func (b *Board) At(i, j int) *Position {
	return &b.Positions[i*b.Size+j]
}

func (b *Board) Distance(p1, p2 Position) int {
	d1 := b.coordDistance(p1.Col, p2.Col)
	d2 := b.coordDistance(p1.Row, p2.Row)
	return d1*d1 + d2*d2
}

func (b *Board) coordDistance(x1, x2 int) int {
	d1 := essentials.AbsInt(x1 - x2)
	d2 := b.Size - d1
	if d2 < d1 {
		return d2
	} else {
		return d1
	}
}

func (b *Board) Nearness() int {
	if b.nearnessCache != 0 {
		return b.nearnessCache
	}
	var res int
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			a1 := *b.At(i, j)
			b1 := Position{Row: i, Col: j}
			for k := 0; k < b.Size; k++ {
				for l := 0; l < b.Size; l++ {
					a2 := *b.At(k, l)
					b2 := Position{Row: k, Col: l}
					if b1.LessThan(b2) {
						res += b.Distance(a1, a2) * b.Distance(b1, b2)
					}
				}
			}
		}
	}
	b.nearnessCache = res
	return res
}

func (b *Board) NormNearness() int {
	table := map[int]int{
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

func (b *Board) String() string {
	rows := make([]string, 0, b.Size)
	for i := 0; i < b.Size; i++ {
		cols := make([]string, 0, b.Size)
		for j := 0; j < b.Size; j++ {
			cols = append(cols, b.At(i, j).String())
		}
		rows = append(rows, "("+strings.Join(cols, ", ")+")")
	}
	return strings.Join(rows, ",\n")
}
