package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Block struct {
	BaseSprite
}

func NewBlock(images []*ebiten.Image) *Block {
	block := new(Block)
	block.Images = images
	block.ImageNum = len(images)
	return block
}
