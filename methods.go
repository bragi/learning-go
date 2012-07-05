package main

import("fmt")


type Game struct {
  name string
  finished bool
}

func (game *Game) state() (state string) {
  if game.finished {
    state = "finished"
    } else {
      state = "started"
    }
    return
}

func (game *Game) logStatus() {
  fmt.Println(game.name)
  fmt.Println(game.state())
}

func (game *Game) finish() {
  game.finished = true
}

func main() {
  game := &Game{name: "New game", finished: false}
  game.logStatus()
  game.finish()
  game.logStatus()
}