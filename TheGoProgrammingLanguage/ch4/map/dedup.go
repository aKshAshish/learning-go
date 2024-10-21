package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dedup: scans the console for inputs and prints it back if it has not seen the input earlier.
func Dedup() {
	// seen, scts as a set
	seen := make(map[string]bool)
	// creates a new scanner to get the input.
	input := bufio.NewScanner(os.Stdin)

	// Read line by line and echo back if line has not been seen.
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
