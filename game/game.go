package game

import (
	"image/color"
	"tiny-side-scroll/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *sprite.Player
	Blocks []*sprite.Block
	// Field *field.Field
}

const (
	screenWidth  = 320
	screenHeight = 240

	charWidth  = 16
	charHeight = 16
)

func (g *Game) Init() {
	// g.Field = field.NewField(field.Field_data_1)
	g.Player = sprite.NewPlayer()
	g.Player.Position.X = 160
	g.Player.Position.Y = 50

	// ブロック
	// 床
	for x := 0; x < 640; x += 17 {
		block := sprite.NewBlock()
		block.Position.X = x
		block.Position.Y = 204
		g.Blocks = append(g.Blocks, block)
	}

	// 左の壁
	for y := 0; y < 200; y += 17 {
		block := sprite.NewBlock()
		block.Position.X = 0
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// 右の壁
	for y := 0; y < 200; y += 17 {
		block := sprite.NewBlock()
		block.Position.X = 629
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// 第2床
	for x := 8 * 17; x < 17*13; x += 17 {
		block := sprite.NewBlock()
		block.Position.X = x
		block.Position.Y = 115
		g.Blocks = append(g.Blocks, block)
	}

	block1 := sprite.NewBlock()
	block1.Position.X = 60
	block1.Position.Y = 165
	g.Blocks = append(g.Blocks, block1)

	block2 := sprite.NewBlock()
	block2.Position.X = 95
	block2.Position.Y = 135
	g.Blocks = append(g.Blocks, block2)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{uint8(155), uint8(188), uint8(15), 0xff})

	// g.Field.Player.Move(g.Field.Sprites)
	// g.Field.Player.Action()
	// g.Field.Player.PlayerJavelins.Move(g.Field.Player.ViewPort)

	// g.Field.Player.DrawImage(screen)
	// for _, j := range g.Field.Player.PlayerJavelins {
	// 	j.DrawImage(screen, g.Field.Player.ViewPort)
	// }
	// g.Field.DrawImage(screen, g.Field.Player.ViewPort)

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
