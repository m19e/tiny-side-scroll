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

func (g *Game) Init() {
	g.Field, g.Player = field.NewField(field.Field_data_1)
	g.Camera = &camera.Camera{}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{uint8(155), uint8(188), uint8(15), 0xff})

	g.Player.Move(g.Field.Sprites, g.Camera)
	g.Player.Action(g.Camera)
	g.Player.Javelins.Move(g.Camera)

	g.Player.DrawImage(screen, nil)
	for _, j := range g.Player.Javelins {
		j.DrawImage(screen, g.Camera)
	}
	g.Field.DrawImage(screen, g.Camera)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
