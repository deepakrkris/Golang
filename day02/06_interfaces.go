package main

import (
	"fmt"
	"math"
)

type SqrtInterface interface {
	get_sqrt() float64
}

type Vertex struct {
	x float64
	y float64
}

func (v Vertex) get_sqrt() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

type Float struct {
	x float64
}

func (f Float) get_sqrt() float64 {
	return math.Sqrt(f.x)
}

func sqrt_processor(num SqrtInterface) float64 {
	return num.get_sqrt()
}

func main() {
	v := Vertex {
		x: 8,
		y: 4,
	}

	fmt.Println(sqrt_processor(v))
	fmt.Println(SqrtInterface(v).get_sqrt())

	f := Float {
		x: 80.00,
	}

	fmt.Println(sqrt_processor(f))
}
