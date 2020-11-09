package sprite

import (
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	javelin_img = `-----+-
+++++++
-----+-`
	javelinSpeed = 4
	screenWidth  = 320
	screenHeight = 240
)

type Javelins []*Javelin

var (
	javelinImg *ebiten.Image
)

func createImageFromString(charString string, img *image.RGBA) {
	width := img.Rect.Size().X
	for indexY, line := range strings.Split(charString, "\n") {
		for indexX, str := range line {
			pos := 4*indexY*width + 4*indexX
			if string(str) == "+" {
				img.Pix[pos] = uint8(15)   // R
				img.Pix[pos+1] = uint8(56) // G
				img.Pix[pos+2] = uint8(15) // B
				img.Pix[pos+3] = 0xff      // A
			} else {
				img.Pix[pos] = uint8(155)   // R
				img.Pix[pos+1] = uint8(188) // G
				img.Pix[pos+2] = uint8(15)  // B
				img.Pix[pos+3] = 0          // A
			}
		}
	}
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, 7, 3))
	createImageFromString(javelin_img, tmpImage)
	javelinImg = ebiten.NewImage(7, 3)
	javelinImg.ReplacePixels(tmpImage.Pix)
}

type Javelin struct {
	BaseSprite
}

func NewJavelin(pos position) *Javelin {
	javelin := new(Javelin)
	javelin.Images = []*ebiten.Image{
		javelinImg,
	}
	javelin.ImageNum = len(javelin.Images)
	javelin.Position = pos
	return javelin
}

func (js *Javelins) Move(viewport position) {
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
