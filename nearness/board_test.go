package nearness

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
			t.Errorf("expected nearness %d but got %d", l2, l1)
		}
	}
}

func BenchmarkBoardMutation(b *testing.B) {
	board := NewBoard(30).Shuffle()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.Mutate()
		board.Nearness()
	}
}
