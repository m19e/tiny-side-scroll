package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m19e/tiny-side-scroll/field"
)

type Game struct {
	Field *field.Field
}

const (
	screenWidth  = 320
	screenHeight = 240

	charWidth  = 16
	charHeight = 16
)

func (g *Game) Init() {
	g.Field = field.NewField(field.Field_data_1)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{uint8(155), uint8(188), uint8(15), 0xff})

	g.Field.Player.Move(g.Field.Sprites)
	g.Field.Player.Action()
	g.Field.Player.PlayerJavelins.Move(g.Field.Player.ViewPort)

	g.Field.Player.DrawImage(screen)
	for _, j := range g.Field.Player.PlayerJavelins {
		j.DrawImage(screen, g.Field.Player.ViewPort)
	}
	g.Field.DrawImage(screen, g.Field.Player.ViewPort)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
