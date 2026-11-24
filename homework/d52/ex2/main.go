package main

import "fmt"

func main() {
	var count int
	var rating float64
	var note string
	var paid bool

	fmt.Printf("int: %v\nfloat64: %v\nstring: %q\nbool: %v\n", count, rating, note, paid)

	// Panic when writing to nil map: panic: assignment to entry in nil map
	fees := make(map[string]int)
	fees["Bole"] = 3000
	fmt.Printf("fees: %v\n", fees)
}
