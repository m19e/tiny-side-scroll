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

func (s *BaseSprite) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Position.X+camera.X), float64(s.Position.Y+camera.Y))
	screen.DrawImage(s.currentImage(), op)
}

func (s *BaseSprite) GetCoordinates() (int, int, int, int) {
	w, h := s.currentImage().Size()
	return s.Position.X, s.Position.Y, w, h
}

func (s *BaseSprite) detectCollisions(object Sprite, dx, dy *int, camera *camera.Camera) *CollideMap {
	var cm CollideMap
	x := s.Position.X
	y := s.Position.Y
	img := s.currentImage()
	w, h := img.Size()

	x1, y1, w1, h1 := object.GetCoordinates()

	x1 += camera.X
	y1 += camera.Y + 1 // +1 for land correctly

	overlappedX := isOverlap(x, x+w, x1, x1+w1)
	overlappedY := isOverlap(y, y+h, y1, y1+h1)

	if overlappedY {
		if *dx < 0 && x+*dx <= x1+w1 && x+w+*dx >= x1 {
			cm.Left = true
		} else if *dx > 0 && x+w+*dx >= x1 && x+*dx <= x1+w1 {
			cm.Right = true
		}
	}
	if overlappedX {
		if *dy < 0 && y+*dy <= y1+w1 && y+h+*dy >= y1 {
			cm.Top = true
		} else if *dy > 0 && y+h+*dy >= y1 && y+*dy <= y1+h1 {
			cm.Bottom = true
		}
	}

	return &cm
}

func (s *BaseSprite) IsCollide(object Sprite, dx, dy *int, camera *camera.Camera) {
	logrus.Info("overwrite this method.")
}

func (s *BaseSprite) Collision(object Sprite, dx, dy *int, cm *CollideMap) {
	logrus.Info("overwrite this method.")
}
