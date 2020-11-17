package game

import (
	"image/color"
	"tiny-side-scroll/camera"
	"tiny-side-scroll/field"
	"tiny-side-scroll/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Field        *field.Field
	Camera       *camera.Camera
	Player       *sprite.Player
}

const (
	screenWidth  = 320
	screenHeight = 240

	charWidth  = 16
	charHeight = 16
)

func (g *Game) Init() {
	g.Field, g.Player = field.NewField(field.Field_data_1)
	g.Camera = &camera.Camera{}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{uint8(155), uint8(188), uint8(15), 0xff})

	g.Player.Move(g.Field.Sprites)
	g.Player.Action()
	g.Player.Javelins.Move(g.Player.ViewPort)

	g.Player.DrawImage(screen, sprite.Position{})
	for _, j := range g.Player.Javelins {
		j.DrawImage(screen, g.Player.ViewPort)
	}
	g.Field.DrawImage(screen, g.Player.ViewPort)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
