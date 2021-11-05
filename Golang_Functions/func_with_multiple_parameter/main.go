package main

import "fmt"

func newMessage(msg string, idx int) {
	fmt.Println(msg, "The value of index is :", idx)
}

func main() {
	fmt.Println("\n************ Func With Multiple Parameter *********")
	for i := 0; i < 5; i++ {
		newMessage("Hello GO", i)
	}
}
