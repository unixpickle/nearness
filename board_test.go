package main

import "testing"

func TestBoardMutation(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := NewBoard(30).Shuffle()
		b.Nearness()
		b.Mutate()
		l1 := b.Nearness()
		b.nearnessCache = 0
		l2 := b.Nearness()
		if l1 != l2 {
			t.Errorf("expected nearness %f but got %f", l2, l1)
		}
	}
}
