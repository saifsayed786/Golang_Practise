package main

import "fmt"

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
	fmt.Println("\n**************** Func as methods ****************")
	res := info{
		firstName: "saif",
		lastname:  "sayed",
	}
	res.information()
	fmt.Println(res)

}
