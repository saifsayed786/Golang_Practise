package main

import (
	"fmt"
)

// func with one parameter
func sayMessage(msg string) {
	fmt.Println(msg)

}

// func with multiple parameter
func newMessage(msg string, idx int) {
	fmt.Println(msg, "The value of index is :", idx)
}

// func with multiple parameter but same types
func sayGreetings(greetings, name string) {
	fmt.Println(greetings, name)

}

//variadic function with return type sum of list

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

// func as method
type info struct {
	firstName string
	lastname  string
	age       int
	marks     int
}

func (i *info) information() {
	fmt.Println(i.firstName + " " + i.lastname)
	i.age = 23
	i.marks = 50

}

func main() {
	fmt.Println("************ Func With One Parameter *************** ")
	sayMessage("Hi Go")

	fmt.Println("\n************ Func With Multiple Parameter *********")
	for i := 0; i < 5; i++ {
		newMessage("Hello GO", i)
	}
	sayGreetings("Hello", "saif")

	fmt.Println("\n************ Variadic Func with Multiple Return Type **************")
	sq, cb := operation(1, 2, 3, 4)
	fmt.Println("Operation Result", sq, cb)
	fmt.Println("\n**************** Func as methods ****************")
	res := info{
		firstName: "saif",
		lastname:  "sayed",
	}
	res.information()
	fmt.Println(res)

}
