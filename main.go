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
	for step := 0; true; step++ {
		for i := MinSize; i <= MaxSize; i++ {
			solutions[i] = Search(solutions[i], 0)
		}
		log.Printf("step %d: score=%f", step, NormalizedScore(solutions))
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

func NormalizedScore(solutions map[int]*Board) float64 {
	rawScores, err := BestRawScores()
	essentials.Must(err)
	var sum float64
	for size, b := range solutions {
		sum += float64(rawScores[size]) / float64(b.NormNearness())
	}
	return sum
}
