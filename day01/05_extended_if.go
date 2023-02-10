package main

import (
	"fmt"
	"math"
)

func pow(n , p , lim float64) float64 {
	if v := math.Pow(n, p) ; v < lim {
		return v
	}
	return lim
}

func main () {
   fmt.Println(pow(3, 2, 10))
   fmt.Println(pow(4, 2, 10))
}
