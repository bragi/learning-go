package main

import 	"code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for x := range rows {
		rows[x] = make([]uint8, dx)
		for y := range rows[x] {
			rows[x][y] = uint8(x)*uint8(y)
		}
	}
	return rows
}

func main() {
	pic.Show(Pic)
}