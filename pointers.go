package main

import(
  "fmt"
)

type Game struct {
  name string
}

func main() {
  gameA := Game{"a"}
  gameB := gameA
  gameAp := &gameA
  fmt.Println(gameA, gameB, gameAp)
  gameB.name = "b"
  gameAp.name = "p"
  fmt.Println(gameA, gameB, gameAp)
}