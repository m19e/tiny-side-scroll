package main

import (
	"log"

	"tiny-side-scroll/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	W = 320
	H = 240
)

func main() {
	game := &game.Game{
		ScreenWidth:  W,
		ScreenHeight: H,
	}
	game.Init()

	ebiten.SetWindowSize(W, H)
	ebiten.SetWindowTitle("tiny-side-scroll")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
