package main

import (
	"image"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m19e/tiny-side-scroll/sprite"
)

var player_anim0 = `-----++-++-----
----+--+--+----
---+-+---+-+---
---+-------+---
--+-+--+--+-+--
--+---+-+---+--
-+--++---++--+-
-+-+-+---+-+-+-
+--+-+---+--+-+
-+-+-------+-+-
--+-+++++++-+--
-----+-+-+-----
-----+-+-+-----
----+-----+----
---+-+-+-+-+---
---++-+-+-++---
-----+---+-----
-----+---+-----`

var player_anim1 = `-----++-++-----
----+--+--+----
---+-+---+-+---
---+-------+---
--+-+--+--+-+--
--+---+-+---+--
-+--++---++--+-
-+-+-+---+-+-+-
+--+-+---+--+-+
-+-+-------+-+-
--+-+++++++-+--
-----+-+-+-----
-----+-+-+-----
----+-----+----
---+-+-+-+-+---
---++-+-+-++---
-----+---+-----
---------+-----`

var player_anim2 = `-----++-++-----
----+--+--+----
---+-+---+-+---
---+-------+---
--+-+--+--+-+--
--+---+-+---+--
-+--++---++--+-
-+-+-+---+-+-+-
+--+-+---+--+-+
-+-+-------+-+-
--+-+++++++-+--
-----+-+-+-----
-----+-+-+-----
----+-----+----
---+-+-+-+-+---
---++-+-+-++---
-----+---+-----
-----+---------`

var block_img = `+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++
+++++++++++++++`

func createImageFromString(charString string, img *image.RGBA) {
	for indexY, line := range strings.Split(charString, "\n") {
		for indexX, str := range line {
			pos := 4*indexY*charWidth + 4*indexX
			if string(str) == "+" {
				img.Pix[pos] = uint8(15)   // R
				img.Pix[pos+1] = uint8(56) // G
				img.Pix[pos+2] = uint8(15) // B
				img.Pix[pos+3] = 0xff      // A
			} else {
				img.Pix[pos] = uint8(155)   // R
				img.Pix[pos+1] = uint8(188) // G
				img.Pix[pos+2] = uint8(15)  // B
				img.Pix[pos+3] = 0xff       // A
			}
		}
	}
}

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

var (
	charWidth   = 15
	charHeight  = 18
	tmpImage    *image.RGBA
	playerAnim0 *ebiten.Image
	playerAnim1 *ebiten.Image
	playerAnim2 *ebiten.Image
	blockImg    *ebiten.Image
)

type Game struct {
	Player *sprite.Player
	Blocks []*sprite.Block
}

func (g *Game) Init() {
	tmpImage = image.NewRGBA(image.Rect(0, 0, charWidth, charHeight))

	createImageFromString(player_anim0, tmpImage)
	playerAnim0 = ebiten.NewImage(charWidth, charHeight)
	playerAnim0.ReplacePixels(tmpImage.Pix)

	createImageFromString(player_anim1, tmpImage)
	playerAnim1 = ebiten.NewImage(charWidth, charHeight)
	playerAnim1.ReplacePixels(tmpImage.Pix)

	createImageFromString(player_anim2, tmpImage)
	playerAnim2 = ebiten.NewImage(charWidth, charHeight)
	playerAnim2.ReplacePixels(tmpImage.Pix)

	createImageFromString(block_img, tmpImage)
	blockImg = ebiten.NewImage(charWidth, charHeight)
	blockImg.ReplacePixels(tmpImage.Pix)

	// プレイヤー
	images := []*ebiten.Image{
		playerAnim0,
		playerAnim1,
		playerAnim2,
	}
	g.Player = sprite.NewPlayer(images)
	g.Player.Position.X = 10
	g.Player.Position.Y = 10

	// ブロック
	block1 := sprite.NewBlock([]*ebiten.Image{blockImg})
	block1.Position.X = 100
	block1.Position.Y = 50
	block2 := sprite.NewBlock([]*ebiten.Image{blockImg})
	block2.Position.X = 200
	block2.Position.Y = 100

	g.Blocks = []*sprite.Block{
		block1,
		block2,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Move([]sprite.Sprite{g.Blocks[0], g.Blocks[1]})

	g.Player.DrawImage(screen)
	for _, block := range g.Blocks {
		block.DrawImage(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}
	game.Init()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("tiny-side-scroll")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
