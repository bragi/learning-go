package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("Sqrt can not accept negative values: %f", e)
}

func Sqrt(f float64) (float64, error) {
  if f < 0 {
    e := ErrNegativeSqrt(f)
    return 0, e
  }
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
