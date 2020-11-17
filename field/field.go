package field

import (
	"strings"

	"tiny-side-scroll/camera"
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
}

func NewField(fieldData string) (*Field, *sprite.Player) {
	player := new(sprite.Player)
	field := new(Field)

	for indexY, line := range strings.Split(fieldData, "\n") {
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
		}
	}

	return field, player
}

func (f *Field) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	for _, sprite := range f.Sprites {
		sprite.DrawImage(screen, camera)
	}
}
