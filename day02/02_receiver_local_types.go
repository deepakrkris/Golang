package main

import (
	"fmt"
	"math"
)

type Float64 float64

func (f Float64) Abs() Float64 {
	if f < 0 {
		return Float64(float64(-f))
	}
	return f
}

func main() {
	var f Float64 = -math.Sqrt2
	fmt.Println(f.Abs())
}
