package sprite

import (
	"image"

	"tiny-side-scroll/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	block_img = `++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++`
	blockWidth  = 16
	blockHeight = 16
)

var (
	blockImg *ebiten.Image
)

type Block struct {
	BaseSprite
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, blockWidth, blockHeight))
	utils.CreateImageFromString(block_img, tmpImage)
	blockImg = ebiten.NewImage(blockWidth, blockHeight)
	blockImg.ReplacePixels(tmpImage.Pix)

}

func NewBlock() *Block {
	block := new(Block)
	block.Images = []*ebiten.Image{
		blockImg,
	}
	block.ImageNum = len(block.Images)
	return block
}
