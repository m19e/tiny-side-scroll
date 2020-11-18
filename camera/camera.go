package camera

type Camera struct {
	X int
	Y int

	Width     int
	Height    int
	MaxWidth  int
	MaxHeight int

	//XLeftLimit  int
	//XRightLimit int
	//YUpperLimit int
	//YLowerLimit int
}

func (c *Camera) SimpleMove(x, y int) {
	c.X = (c.Width / 2) - x
	c.Y = (c.Height / 2) - y
}

func (c *Camera) Move(x, y int) {
	maxXOffset := -(c.MaxWidth - c.Width)
	maxYOffset := -(c.MaxHeight - c.Height)

	restX := (c.Width / 2) - x
	restY := (c.Height / 2) - y

	if restX > 0 {
		c.X = 0
	} else if restX < maxXOffset {
		c.X = maxXOffset
	} else {
		c.X = restX
	}

	if restY > 0 {
		c.Y = 0
	} else if restY < maxYOffset {
		c.Y = maxYOffset
	} else {
		c.Y = restY
	}
}
