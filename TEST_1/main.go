package main

import (
	"fmt"
)

func main() {
	// Prompt user for input
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	// Generate triangular numbers
	var output string
	for i := 1; i <= n; i++ {
		output += fmt.Sprintf("%d-", (i*(i+1))/2)
	}

	// Print the output
	fmt.Println(output[:len(output)-1])
}
