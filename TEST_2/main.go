package main

import (
	"fmt"
	"sort"
)

func main() {
	// Prompt user for number of players and scores
	var n int
	fmt.Print("Enter number of players: ")
	fmt.Scan(&n)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter score for player %d: ", i+1)
		fmt.Scan(&scores[i])
	}

	// Sort scores in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	// Calculate dense rankings
	var rankings = make(map[int]int)
	currentRank := 1
	for i, score := range scores {
		if i > 0 && scores[i-1] != score {
			currentRank = i + 1
		}
		rankings[score] = currentRank
	}

	// Prompt user for GITS scores
	var m int
	fmt.Print("Enter number of GITS scores: ")
	fmt.Scan(&m)

	var gitsScores []int
	for i := 0; i < m; i++ {
		var score int
		fmt.Printf("Enter score for GITS %d: ", i+1)
		fmt.Scan(&score)
		gitsScores = append(gitsScores, score)
	}

	// Print GITS rankings
	for _, score := range gitsScores {
		fmt.Printf("GITS rank: %d\n", rankings[score])
	}
}
