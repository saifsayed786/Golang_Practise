package main

import "fmt"

// Function with int as return type
func addt(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	// Accepting return value in varaible
	sum := addt(20, 30)
	fmt.Println(sum)
}
