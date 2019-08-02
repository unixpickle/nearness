package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

const NumImprovers = 4

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
	for step := 0; step < 10000; step++ {
		var wg sync.WaitGroup
		for _, imp := range solutions {
			wg.Add(1)
			go func(imp *Improver) {
				defer wg.Done()
				imp.Step()
			}(imp)
		}
		wg.Wait()
		log.Printf("loss30=%d loss6=%d", solutions[30].BestBoard().NormNearness(),
			solutions[6].BestBoard().NormNearness())
		if step%10 == 0 {
			log.Println("saving solution set")
			SaveSolutions(solutions)
		}
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
