package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m19e/tiny-side-scroll/game"
)

func main() {
	game := &game.Game{}
	game.Init()

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("tiny-side-scroll")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
