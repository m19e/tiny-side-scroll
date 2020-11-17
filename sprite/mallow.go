package sprite

import (
	"image"

	"tiny-side-scroll/camera"
	"tiny-side-scroll/utils"

	"github.com/hajimehoshi/ebiten/v2"
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
	utils.CreateImageFromString(mallow_img, tmpImage, utils.Green)
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

func (m *Mallow) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	if m.Alive {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(m.Position.X+camera.X), float64(m.Position.Y+camera.Y))
		screen.DrawImage(m.currentImage(), op)
	}
}
