package sprite

import (
	"tiny-side-scroll/camera"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

type Sprite interface {
	GetCoordinates() (int, int, int, int)
	DrawImage(*ebiten.Image, *camera.Camera)
	Collision(Sprite, *int, *int)
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

func (s *BaseSprite) Intersect(object Sprite) bool {
	ax, ay, aw, ah := s.GetCoordinates()
	bx, by, bw, bh := object.GetCoordinates()

	// a left-top < b right-bottom && a right-bottom > b left-top
	return (ax < bx+bw && ay < by+bh) && (ax+aw > bx && ay+ah > by)
}

func (s *BaseSprite) Width() int {
	w, _ := s.currentImage().Size()
	return w
}

func (s *BaseSprite) Height() int {
	_, h := s.currentImage().Size()
	return h
}

func (s *BaseSprite) IsCollide(object Sprite, dx, dy *int, camera *camera.Camera) {
	logrus.Info("overwrite this method.")
}

func (s *BaseSprite) Collision(object Sprite, dx, dy *int) {
	logrus.Info("overwrite this method.")
}
