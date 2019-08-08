package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/unixpickle/essentials"
)

const (
	MinSize = 6
	MaxSize = 30
)

func main() {
	solutions := map[int]*Board{}
	for i := MinSize; i <= MaxSize; i++ {
		solutions[i] = NewBoard(i).Shuffle()
	}
	if data, err := ioutil.ReadFile("solutions.txt"); err == nil {
		boards, err := ParseBoards(string(data))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("using saved boards...")
		for _, b := range boards {
			solutions[b.Size] = b
		}
	}
	exploiter := NewExploiter(MinSize, MaxSize)
	for step := 0; true; step++ {
		size := exploiter.Sample()
		oldScore := NormalizedScore(solutions[size])
		solutions[size] = Search(solutions[size], 0)
		newScore := NormalizedScore(solutions[size])
		exploiter.GotUtility(size, newScore-oldScore)
		log.Printf("step %d: score=%f delta=%f size=%d", step, TotalNormalizedScore(solutions),
			newScore-oldScore, size)
		SaveSolutions(solutions)
	}
}

func SaveSolutions(solutions map[int]*Board) {
	strs := make([]string, 0, len(solutions))
	for i := MinSize; i <= MaxSize; i++ {
		strs = append(strs, solutions[i].String())
	}
	data := []byte(strings.Join(strs, ";\n"))
	ioutil.WriteFile("solutions.txt", data, 0755)
}

func TotalNormalizedScore(solutions map[int]*Board) float64 {
	var sum float64
	for _, b := range solutions {
		sum += NormalizedScore(b)
	}
	return sum
}

func NormalizedScore(b *Board) float64 {
	rawScores, err := BestRawScores()
	essentials.Must(err)
	return float64(rawScores[b.Size]) / float64(b.NormNearness())
}
