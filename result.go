package main

import "fmt"

type Result struct {
	TotalScore int
	Scores     map[string]int // key: player name, value: total score of that player
	MatchesWon map[string]int // key: player name, value: matches won by that player
}

func (r *Result) Print() {
	fmt.Println("TotalScore:", r.TotalScore)
	fmt.Println()

	for k, v := range r.Scores {
		fmt.Printf("Score of %v: %v\n", k, v)
	}

	fmt.Println("\nPeople that won at least one game:")
	for k, v := range r.MatchesWon {
		fmt.Printf("matches won by %v -> %v\n", k, v)
	}
}
