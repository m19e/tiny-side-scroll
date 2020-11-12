package main

import (
	"log"

	"tiny-side-scroll/game"

	"github.com/hajimehoshi/ebiten/v2"
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
