package utils

import (
	"image"
	"image/color"
	"strings"
)

var (
	Green = &color.RGBA{uint8(15), uint8(56), uint8(15), 0xff}
)

func CreateImageFromString(charString string, img *image.RGBA) {
	width := img.Rect.Size().X
	for indexY, line := range strings.Split(charString, "\n") {
		for indexX, str := range line {
			pos := 4*indexY*width + 4*indexX
			if string(str) == "+" {
				img.Pix[pos] = uint8(15)   // R
				img.Pix[pos+1] = uint8(56) // G
				img.Pix[pos+2] = uint8(15) // B
				img.Pix[pos+3] = 0xff      // A
			} else {
				img.Pix[pos] = 0   // R
				img.Pix[pos+1] = 0 // G
				img.Pix[pos+2] = 0 // B
				img.Pix[pos+3] = 0 // A
			}
		}
	}
}
