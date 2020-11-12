package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m19e/tiny-side-scroll/sprite"
)

type Game struct {
	Player *sprite.Player
	Blocks []*sprite.Block
}

const (
	screenWidth  = 320
	screenHeight = 240

	charWidth  = 16
	charHeight = 16
)

func (g *Game) Init() {
	g.Player = sprite.NewPlayer()
	g.Player.Position.X = 160
	g.Player.Position.Y = 50

	for x := 0; x < 640; x += (charWidth + 1) {
		block := sprite.NewBlock()
		block.Position.X = x
		block.Position.Y = 204
		g.Blocks = append(g.Blocks, block)
	}
	// floor above
	for x := 9 * (charWidth + 1); x < 13*(charWidth+1); x += (charWidth + 1) {
		block := sprite.NewBlock()
		block.Position.X = x
		block.Position.Y = 115
		g.Blocks = append(g.Blocks, block)
	}

	// left wall
	for y := 0; y < 200; y += (charHeight + 1) {
		block := sprite.NewBlock()
		block.Position.X = 0
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// right wall
	for y := 0; y < 200; y += (charHeight + 1) {
		block := sprite.NewBlock()
		block.Position.X = 629
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	block1 := sprite.NewBlock()
	block1.Position.X = 4 * (charWidth + 1)
	block1.Position.Y = 160
	g.Blocks = append(g.Blocks, block1)

	block2 := sprite.NewBlock()
	block2.Position.X = 6 * (charWidth + 1)
	block2.Position.Y = 140
	g.Blocks = append(g.Blocks, block2)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{uint8(155), uint8(188), uint8(15), 0xff})

	sprites := []sprite.Sprite{}
	for _, b := range g.Blocks {
		sprites = append(sprites, b)
	}
	g.Player.Move(sprites)
	g.Player.Action()
	g.Player.PlayerJavelins.Move(g.Player.ViewPort)

	g.Player.DrawImage(screen)
	for _, j := range g.Player.PlayerJavelins {
		j.DrawImage(screen, g.Player.ViewPort)
	}
	for _, block := range g.Blocks {
		block.DrawImage(screen, g.Player.ViewPort)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}