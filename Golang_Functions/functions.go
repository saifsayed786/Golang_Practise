package main

import "fmt"

// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

// Function with int as return type
func addt(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	// Passing arguments
	add(20, 30)
	// Accepting return value in varaible
	sum := addt(20, 30)
	fmt.Println(sum)
}
