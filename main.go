package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/unixpickle/nearness/nearness"
)

const (
	MinSize = 6
	MaxSize = 30
)

func main() {
	solutions := map[int]*nearness.Board{}
	for i := MinSize; i <= MaxSize; i++ {
		solutions[i] = nearness.NewBoard(i).Shuffle()
	}
	if data, err := ioutil.ReadFile("solutions.txt"); err == nil {
		boards, err := nearness.ParseBoards(string(data))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("using saved boards...")
		for _, b := range boards {
			solutions[b.Size] = b
		}
	}
	exploiter := nearness.NewExploiter(MinSize, MaxSize)
	for step := 0; true; step++ {
		size := exploiter.Sample()
		oldScore := nearness.NormalizedScore(solutions[size])
		solutions[size] = nearness.Search(solutions[size], 0)
		newScore := nearness.NormalizedScore(solutions[size])
		exploiter.GotUtility(size, newScore-oldScore)
		log.Printf("step %d: score=%f delta=%f size=%d",
			step,
			nearness.TotalNormalizedScore(solutions),
			newScore-oldScore,
			size)
		SaveSolutions(solutions)
	}
}

func SaveSolutions(solutions map[int]*nearness.Board) {
	strs := make([]string, 0, len(solutions))
	for i := MinSize; i <= MaxSize; i++ {
		strs = append(strs, solutions[i].String())
	}
	data := []byte(strings.Join(strs, ";\n"))
	ioutil.WriteFile("solutions.txt", data, 0755)
}
