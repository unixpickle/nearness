package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

func main() {
	solutions := map[int]*Board{}
	for i := 6; i <= 30; i++ {
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
	for step := 0; step < 10000; step++ {
		for _, b := range solutions {
			go ImproveSolution(b)
		}
		log.Printf("loss30=%f loss6=%f", solutions[30].NormNearness(), solutions[6].NormNearness())
		if step%10 == 0 {
			log.Println("saving solution set")
			SaveSolutions(solutions)
		}
	}
}

func ImproveSolution(b *Board) {
	origLoss := b.Nearness()
	b1 := b.Copy()
	for i := 0; i < 1000; i++ {
		for i := 0; i < 1+rand.Intn(3); i++ {
			b1.Mutate()
		}
		if b1.Nearness() < b.Nearness() {
			*b = *b1
		} else {
			*b1 = *b
		}
	}
	if b.Nearness() == origLoss {
		// Let's try to push ourselves out of this global
		// minimum.
		for i := 0; i < 30; i++ {
			b1.Mutate()
		}
		ImproveSolution(b1)
		if b1.Nearness() < origLoss {
			*b = *b1
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
