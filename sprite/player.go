package sprite

import (
	"image"
	"math"

	"tiny-side-scroll/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

func isOverlap(x1, x2, x3, x4 int) bool {
	if x1 <= x4 && x2 >= x3 {
		return true
	}
	return false
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

func (p *Player) Move(objects []Sprite, viewPort *Position) {
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
		p.IsCollide(object, &dx, &dy, viewPort)
	}

	if p.Position.X+dx < xLeftLimit || p.Position.X+dx > xRightLimit {
		viewPort.X -= dx
	} else {
		p.Position.X += dx
	}

	if p.Position.Y+dy < yUpperLimit || p.Position.Y+dy > yLowerLimit {
		viewPort.Y -= dy
	} else {
		p.Position.Y += dy
	}
}

func (p *Player) Action(viewPort *Position) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		pos := Position{
			X: (p.Position.X - viewPort.X) + 8,
			Y: (p.Position.Y - viewPort.Y) + 4,
		}
		javelin := NewJavelin(pos)
		p.Javelins = append(p.Javelins, javelin)
	}
}

func (p *Player) DrawImage(screen *ebiten.Image, _ Position) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.currentImage(), op)
}

func (p *Player) collideBlock(_ *Block, dx, dy *int, cm *CollideMap) {
	if cm.Left || cm.Right {
		*dx = 0
	}
	if cm.Top {
		*dy = 0
	}
	if cm.Bottom {
		*dy = 0
		p.jumping = false
		p.jumpSpeed = 0
	}
}
