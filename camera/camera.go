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

	XCenter int
	YCenter int
}

func (c *Camera) Move(x, y int) {
	maxXOffset := -(c.MaxHeight - c.Width)
	maxYOffset := -(c.MaxWidth - c.Height)

	restX := c.XCenter - x
	restY := c.YCenter - y

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
