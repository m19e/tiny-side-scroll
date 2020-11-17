package field

import (
	"strings"

	"tiny-side-scroll/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 17
	height = 17

	blockMark  = "+"
	playerMark = "P"
	mallowMark = "M"
)

type Field struct {
	Sprites []sprite.Sprite
	Width   int
	Height  int
}

func NewField(fieldData string) (*Field, *sprite.Player) {
	player := new(sprite.Player)
	field := new(Field)

	var countX, countY int

	for indexY, line := range strings.Split(fieldData, "\n") {
		counter := 0
		for indexX, str := range line {
			switch string(str) {
			case blockMark:
				block := sprite.NewBlock()
				block.Position.X = indexX * width
				block.Position.Y = indexY * height
				field.Sprites = append(field.Sprites, block)
			case playerMark:
				player = sprite.NewPlayer()
				player.Position.X = indexX * width
				player.Position.Y = indexY * height
			case mallowMark:
				mallow := sprite.NewMallow()
				mallow.Position.X = indexX * width
				mallow.Position.Y = indexY * height
				field.Sprites = append(field.Sprites, mallow)
			}
			counter++
		}
		if countX < counter {
			countX = counter
		}
		countY++
	}

	field.Width = countX * width
	field.Height = countY * height
	return field, player
}

func (f *Field) DrawImage(screen *ebiten.Image, viewport sprite.Position) {
	for _, sprite := range f.Sprites {
		sprite.DrawImage(screen, viewport)
	}
}
