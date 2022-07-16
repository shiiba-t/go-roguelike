package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//Game holds all data the entire game will need.
type Game struct {
}

//NewGame creates a new Game Object and initializes the data
//This is a pretty solid refactor candidate for later
func NewGame() *Game {
	g := &Game{}
	return g
}

//Update is called each tic.
func (g *Game) Update() error {
	return nil
}

//Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
}

//Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
