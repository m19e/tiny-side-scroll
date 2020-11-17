package sprite

import (
	"tiny-side-scroll/camera"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Sprite interface {
	GetCoordinates() (int, int, int, int)
	DrawImage(*ebiten.Image, *camera.Camera)
	Collision(Sprite, *int, *int, *CollideMap)
}

type Position struct {
	X int
	Y int
}

type CollideMap struct {
	Left   bool
	Right  bool
	Top    bool
	Bottom bool
}

func (cm *CollideMap) HasCollision() bool {
	return cm.Left || cm.Right || cm.Top || cm.Bottom
}

type BaseSprite struct {
	Images     []*ebiten.Image // Image array for Animation sprite
	ImageNum   int             // Amount of all images
	CurrentNum int             // Current number display image
	Position   Position        // Current position display image
	count      int             // Counter for frame number
}

func NewSprite(images []*ebiten.Image) *BaseSprite {
	return &BaseSprite{
		Images:   images,
		ImageNum: len(images),
	}
}

func (s *BaseSprite) currentImage() *ebiten.Image {
	if s.count > 5 {
		s.count = 0
		s.CurrentNum++
		s.CurrentNum %= s.ImageNum
	}
	return s.Images[s.CurrentNum]
}

func (s *BaseSprite) DrawImage(screen *ebiten.Image, viewPort Position) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Position.X+viewPort.X), float64(s.Position.Y+viewPort.Y))
	screen.DrawImage(s.currentImage(), op)
}

func (s *BaseSprite) GetCoordinates() (int, int, int, int) {
	w, h := s.currentImage().Size()
	return s.Position.X, s.Position.Y, w, h
}

func (s *BaseSprite) Collision(object Sprite, dx, dy *int, cm *CollideMap) {
	logrus.Info("overwrite this method.")
}
