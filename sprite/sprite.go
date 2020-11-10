package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Sprite interface {
	GetCordinates() (int, int, int, int)
}

type Position struct {
	X int
	Y int
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

func (s *BaseSprite) GetCordinates() (int, int, int, int) {
	w, h := s.currentImage().Size()
	return s.Position.X, s.Position.Y, w, h
}
