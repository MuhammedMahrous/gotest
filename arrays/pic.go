package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8
	for y := 0; y < dy; y++ {
		var slice []uint8
		for x := 0 ; x < dx; x++ {
			slice = append(slice, uint8(x*x))
		}
		s = append(s, slice)
	}

	return s
}

func main() {
	pic.Show(Pic)
}
