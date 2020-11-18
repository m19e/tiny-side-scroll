package sprite

import (
	"image"
	"math"

	"tiny-side-scroll/camera"
	"tiny-side-scroll/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
)

const (
	xLeftLimit  = 16 * 2
	xRightLimit = 320 - (16 * 7)
	yUpperLimit = 16 * 2
	yLowerLimit = 240 - (16 * 2)

	charWidth  = 16
	charHeight = 16

	player_anim0 = `-----++--++-----
----+--++--+----
---+-+----+-+---
--+-+--++--+-+--
--+---+--+---+--
-+--++----++--+-
-+-+-+----+-+-+-
+--+-+----+--+-+
-+-+--------+-+-
--+-++++++++-+--
------++++------
-----+-++-+-----
-----+-++-+-----
----+-+--+-+----
----++-++-++----
------+--+------`

	player_anim1 = `-----++--++-----
----+--++--+----
---+-+----+-+---
--+-+--++--+-+--
--+---+--+---+--
-+--++----++--+-
-+-+-+----+-+-+-
+--+-+----+--+-+
-+-+--------+-+-
--+-++++++++-+--
------++++------
-----+-++-+-----
-----+-++-+-----
----+-+--+-+----
----++-++-++----
---------+------`

	player_anim2 = `-----++--++-----
----+--++--+----
---+-+----+-+---
--+-+--++--+-+--
--+---+--+---+--
-+--++----++--+-
-+-+-+----+-+-+-
+--+-+----+--+-+
-+-+--------+-+-
--+-++++++++-+--
------++++------
-----+-++-+-----
-----+-++-+-----
----+-+--+-+----
----++-++-++----
------+---------`
)

var (
	playerAnim0 *ebiten.Image
	playerAnim1 *ebiten.Image
	playerAnim2 *ebiten.Image
)

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, charWidth, charHeight))

	utils.CreateImageFromString(player_anim0, tmpImage, utils.Green)
	playerAnim0 = ebiten.NewImage(charWidth, charHeight)
	playerAnim0.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim1, tmpImage, utils.Green)
	playerAnim1 = ebiten.NewImage(charWidth, charHeight)
	playerAnim1.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim2, tmpImage, utils.Green)
	playerAnim2 = ebiten.NewImage(charWidth, charHeight)
	playerAnim2.ReplacePixels(tmpImage.Pix)
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}

type Player struct {
	BaseSprite
	jumping   bool
	jumpSpeed float64
	fallSpeed float64
	Javelins  Javelins
}

func NewPlayer() *Player {
	player := new(Player)
	player.Images = []*ebiten.Image{
		playerAnim0,
		playerAnim1,
		playerAnim2,
	}
	player.ImageNum = len(player.Images)
	player.jumpSpeed = 0
	player.fallSpeed = 0.4
	return player
}

func (p *Player) jump() {
	if !p.jumping {
		p.jumping = true
		p.jumpSpeed = -7
	}
}

func (p *Player) Move(objects []Sprite, camera *camera.Camera) {
	var dx, dy int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx = -2
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx = 2
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.jump()
		p.count++
	}

	if p.jumpSpeed < 5 {
		p.jumpSpeed += p.fallSpeed
	}
	dy = round(p.jumpSpeed)

	for _, object := range objects {
		p.IsCollide(object, &dx, &dy, camera)
	}

	if p.Position.X+dx < xLeftLimit || p.Position.X+dx > xRightLimit {
		camera.X -= dx
	} else {
		p.Position.X += dx
	}

	if p.Position.Y+dy < yUpperLimit || p.Position.Y+dy > yLowerLimit {
		camera.Y -= dy
	} else {
		p.Position.Y += dy
	}
}

func (p *Player) Action() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		pos := Position{
			X: (p.Position.X) + 8,
			Y: (p.Position.Y) + 4,
		}
		javelin := NewJavelin(pos)
		p.Javelins = append(p.Javelins, javelin)
	}
}

func (p *Player) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X+camera.X), float64(p.Position.Y+camera.Y))
	screen.DrawImage(p.currentImage(), op)
}

func (p *Player) Collision(object Sprite, dx, dy int) {
	switch v := object.(type) {
	case *Block:
		p.collideBlock(v, dx, dy)
	case *Mallow:
		p.collideMallow(v)
	default:
		logrus.Warn("unknown type")
	}
}

func (p *Player) collideBlock(b *Block, dx, dy int) {
	if dx > 0 {
		p.Position.X = b.Position.X - p.Width()
	}
	if dx < 0 {
		p.Position.X = b.Position.X + p.Width()
	}
	if dy < 0 {
		p.Position.Y = p.Position.Y + p.Height()
	}
	if dy > 0 {
		p.Position.Y = b.Position.Y - p.Height()
		p.jumping = false
		p.jumpSpeed = 0
	}
}

func (p *Player) collideMallow(m *Mallow) {
	m.Alive = false
}
