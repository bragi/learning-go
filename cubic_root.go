package main

import (
  "fmt"
  "math/cmplx"
)

func Cbrt(x complex128) complex128 {
  
  Step := func(x, z complex128) complex128 {
    return z - (z*z*z -x)/(3*z*z)
  }

  const LIMIT = 1e-5
  var z complex128 = 1
  var newZ complex128 = Step(x, z)
  
  for cmplx.Abs(newZ - z) > LIMIT {
    z = newZ
    newZ = Step(x, z)
  }
  return newZ
}

func main() {
	fmt.Println(Cbrt(2))
  fmt.Println(cmplx.Pow(2, 0.5))
}
