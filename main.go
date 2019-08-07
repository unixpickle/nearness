package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/unixpickle/essentials"
)

const NumImprovers = 64

func main() {
	solutions := map[int]*Improver{}
	for i := 6; i <= 30; i++ {
		solutions[i] = NewImprover(NewBoard(i).Shuffle(), NumImprovers)
	}
	if data, err := ioutil.ReadFile("solutions.txt"); err == nil {
		boards, err := ParseBoards(string(data))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("using saved boards...")
		for _, b := range boards {
			solutions[b.Size] = NewImprover(b, NumImprovers)
		}
	}
	for step := 0; true; step++ {
		var wg sync.WaitGroup
		for _, imp := range solutions {
			wg.Add(1)
			go func(imp *Improver) {
				defer wg.Done()
				imp.Step()
			}(imp)
		}
		wg.Wait()
		log.Printf("loss30=%d score=%f", solutions[30].BestBoard().NormNearness(),
			NormalizedScore(solutions))
		SaveSolutions(solutions)
	}
}

func SaveSolutions(solutions map[int]*Improver) {
	strs := make([]string, 0, len(solutions))
	for _, imp := range solutions {
		strs = append(strs, imp.BestBoard().String())
	}
	data := []byte(strings.Join(strs, ";\n"))
	ioutil.WriteFile("solutions.txt", data, 0755)
}

func NormalizedScore(solutions map[int]*Improver) float64 {
	rawScores, err := BestRawScores()
	essentials.Must(err)
	var sum float64
	for size, improver := range solutions {
		sum += float64(rawScores[size]) / float64(improver.BestBoard().NormNearness())
	}
	return sum
}
