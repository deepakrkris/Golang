package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][] uint8, dy)

	for y := range make([]int, dy)  {
		r := make([] uint8, dx)
		for x := range make([]int, dx)  {
			r[x] = uint8(x ^ y)
		}
		s[y] = r
	}

	return s
}

func main() {
	pic.Show(Pic)
}