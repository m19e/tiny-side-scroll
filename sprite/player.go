package sprite

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/m19e/tiny-side-scroll/utils"
)

const (
	xLeftLimit  = 16 * 3
	xRightLimit = 320 - (16 * 3)
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

	utils.CreateImageFromString(player_anim0, tmpImage)
	playerAnim0 = ebiten.NewImage(charWidth, charHeight)
	playerAnim0.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim1, tmpImage)
	playerAnim1 = ebiten.NewImage(charWidth, charHeight)
	playerAnim1.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim2, tmpImage)
	playerAnim2 = ebiten.NewImage(charWidth, charHeight)
	playerAnim2.ReplacePixels(tmpImage.Pix)
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}

func isOverlap(x1, x2, x3, x4 int) bool {
	if x1 <= x4 && x2 >= x3 {
		return true
	}
	return false
}

type Player struct {
	BaseSprite
	jumping        bool
	jumpSpeed      float64
	fallSpeed      float64
	ViewPort       Position
	PlayerJavelins Javelins
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

func (p *Player) Move(objects []Sprite) {
	var dx, dy int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx = -1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx = 1
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
		dx, dy = p.IsCollide(dx, dy, object)
	}

	if p.Position.X+dx < xLeftLimit || p.Position.X+dx > xRightLimit {
		p.ViewPort.X -= dx
	} else {
		p.Position.X += dx
	}

	if p.Position.Y+dy < yUpperLimit || p.Position.Y+dy > yLowerLimit {
		p.ViewPort.Y -= dy
	} else {
		p.Position.Y += dy
	}
}

func (p *Player) Action() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		pos := Position{
			X: (p.Position.X - p.ViewPort.X) + 8,
			Y: (p.Position.Y - p.ViewPort.Y) + 4,
		}
		javelin := NewJavelin(pos)
		p.PlayerJavelins = append(p.PlayerJavelins, javelin)
	}
}

func (p *Player) IsCollide(dx, dy int, object Sprite) (int, int) {
	x := p.Position.X
	y := p.Position.Y
	img := p.currentImage()
	w, h := img.Size()

	x1, y1, w1, h1 := object.GetCordinates()

	// 対象オブジェクトは相対座標付与して衝突判定を行う
	x1 += p.ViewPort.X
	y1 += p.ViewPort.Y + 1

	overlappedX := isOverlap(x, x+w, x1, x1+w1)
	overlappedY := isOverlap(y, y+h, y1, y1+h1)

	if overlappedY {
		if dx < 0 && x+dx <= x1+w1 && x+w+dx >= x1 {
			dx = 0
		} else if dx > 0 && x+w+dx >= x1 && x+dx <= x1+w1 {
			dx = 0
		}
	}
	if overlappedX {
		if dy < 0 && y+dy <= y1+h1 && y+h+dy >= y1 {
			dy = 0
		} else if dy > 0 && y+h+dy >= y1 && y+dy <= y1+h1 {
			dy = 0
			p.jumping = false
			p.jumpSpeed = 0
		}
	}

	return dx, dy
}

func (p *Player) DrawImage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.currentImage(), op)
}
