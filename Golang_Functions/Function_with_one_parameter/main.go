package main

import "fmt"

func sayMessage(msg string) {
	fmt.Println(msg)

}
func main() {
	fmt.Println("************ Func With One Parameter *************** ")
	sayMessage("Hi Go")
}
