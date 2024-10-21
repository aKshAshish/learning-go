package main

import "fmt"

func Map() {
	// Initialising a map
	var m map[string]int
	m = map[string]int{
		"one": 1,
		"two": 2,
	}

	mapMake := make(map[string]bool)
	mapMake["ashish"] = true
	fmt.Println(mapMake)

	// Loop over map
	fmt.Print("{\n")
	for k, v := range m {
		fmt.Printf("  %s : %d,\n", k, v)
	}
	fmt.Print("}\n")
}
