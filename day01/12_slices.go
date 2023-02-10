package main

import "fmt"

func main() {
	primes := [6]int {1, 2, 3, 4, 5, 6}

	var s []int = primes[1:4]

	fmt.Println(s)

	s = make([] int, 3)

	fmt.Println(s)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println(s)

	s[3] = 4

	fmt.Println(s)
}
