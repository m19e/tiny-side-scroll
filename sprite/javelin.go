package sprite

import (
	"image"
	"tiny-side-scroll/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	javelin_img = `-----+-
+++++++
-----+-`
	javelinSpeed = 2
	screenWidth  = 320
	screenHeight = 240
)

type Javelins []*Javelin

var (
	javelinImg *ebiten.Image
)

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, 7, 3))
	utils.CreateImageFromString(javelin_img, tmpImage, utils.Green)
	javelinImg = ebiten.NewImage(7, 3)
	javelinImg.ReplacePixels(tmpImage.Pix)
}

type Javelin struct {
	BaseSprite
}

func NewJavelin(pos Position) *Javelin {
	javelin := new(Javelin)
	javelin.Images = []*ebiten.Image{
		javelinImg,
	}
	javelin.ImageNum = len(javelin.Images)
	javelin.Position = pos
	return javelin
}

func (js *Javelins) Move(viewport Position) {
	javelins := *js

	for i := 0; i < len(javelins); i++ {
		j := javelins[i]
		j.Position.X += javelinSpeed

		if j.Position.X > (screenWidth-viewport.X) ||
			j.Position.Y > (screenHeight-viewport.Y) ||
			j.Position.X < 0 ||
			j.Position.Y < 0 {
			javelins = append(javelins[:i], javelins[i+1:]...)
			i--
		}
		*js = javelins
	}
}
