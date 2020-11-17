package game

import (
	"tiny-side-scroll/camera"
	"tiny-side-scroll/field"
	"tiny-side-scroll/sprite"
	"tiny-side-scroll/utils"

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
	g.Camera = &camera.Camera{
		Width:  g.ScreenWidth,
		Height: g.ScreenHeight,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(utils.LightGreen)

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
