package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/m19e/tiny-side-scroll/utils"
)

const (
	mallow_img = `--++++++---
-+------+-
+--------+
++------++
+-++++++-+
+--------+
+--------+
+--------+
+--------+
+--------+
-+------+-
--++++++--`

	mallowWidth  = 10
	mallowHeight = 12
)

var mallowImg *ebiten.Image

type Mallow struct {
	BaseSprite
	Alive bool
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, mallowWidth, mallowHeight))
	utils.CreateImageFromString(mallow_img, tmpImage)
	mallowImg = ebiten.NewImage(mallowWidth, mallowHeight)
	mallowImg.ReplacePixels(tmpImage.Pix)
}

func NewMallow() *Mallow {
	mallow := new(Mallow)
	mallow.Images = []*ebiten.Image{
		mallowImg,
	}
	mallow.ImageNum = len(mallow.Images)
	mallow.Alive = true
	return mallow
}

func (m *Mallow) DrawImage(screen *ebiten.Image, viewPort Position) {
	if m.Alive {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(m.Position.X+viewPort.X), float64(m.Position.Y+viewPort.Y))
		screen.DrawImage(m.currentImage(), op)
	}
}
