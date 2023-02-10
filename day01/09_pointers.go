package main

import (
	"fmt"
)

func incr(x int) int {
	x++
	return x
}

func incr_ref(x *int) int {
	*x++
	return *x
}

func main() {
	x := 10

	fmt.Println("increment by val", x)
	fmt.Println(incr(x))

	fmt.Println("value of x", x)

	fmt.Println("increment by ref", x)
	fmt.Println(incr_ref(&x))

	fmt.Println("value of x", x)
}

