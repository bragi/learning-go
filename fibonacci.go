package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  value := 0
  previousValue := 1
  olderValue :=0
  return func() int {
    olderValue = previousValue
    previousValue = value
    value = olderValue + previousValue
    return value
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
