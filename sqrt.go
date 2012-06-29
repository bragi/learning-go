package main

import "fmt"

func Sqrt(x float64) (z float64) {
	
	Mod := func (value float64) float64 {
		if value < 0 {
			return -value
		}
		return value
	}

	SqrtStep := func(x float64, z float64) float64 {
		return z - (z*z-x)/(2*z)
	}
	
	const LIMIT = 1e-10
	z = 10
	
	newZ := SqrtStep(x, z)
	for Mod(z-newZ) > LIMIT {
		z = newZ
		newZ = SqrtStep(x, z)
	}
	return
}

func main() {
	fmt.Println(Sqrt(7))
}
