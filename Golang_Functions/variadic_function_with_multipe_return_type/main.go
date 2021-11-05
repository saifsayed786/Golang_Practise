package main

import "fmt"

func operation(values ...int) (squared int, cube int) {
	fmt.Println("Actual List", values)
	// squared := 0
	// cube := 0
	for _, v := range values {
		squared += v * v
		cube += v * v * v
	}
	return
	// return squared, cube
}

func main() {
	fmt.Println("\n************ Variadic Func with Multiple Return Type **************")
	sq, cb := operation(1, 2, 3, 4)
	fmt.Println("Operation Result", sq, cb)
}
