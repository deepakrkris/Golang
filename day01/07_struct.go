package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func newPerson() person {
	p := person{name : "Deepak"}
	p.age = 20
	return p
}

func main() {
	p := newPerson()
	fmt.Println(p.name)
}
