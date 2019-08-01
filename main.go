package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	solutions := map[int]*Board{}
	for i := 6; i <= 30; i++ {
		solutions[i] = NewBoard(i).Shuffle()
	}
	for step := 0; step < 10000000; step++ {
		for size, b := range solutions {
			b1 := b.Copy()
			b1.Mutate()
			if b1.Nearness() < b.Nearness() {
				solutions[size] = b1
			}
		}
		if step%10000 == 0 {
			log.Println("saving solution set")
			SaveSolutions(solutions)
		}
	}
}

func SaveSolutions(solutions map[int]*Board) {
	strs := make([]string, 0, len(solutions))
	for _, b := range solutions {
		strs = append(strs, b.String())
	}
	data := []byte(strings.Join(strs, ";\n"))
	ioutil.WriteFile("solutions.txt", data, 0755)
}
