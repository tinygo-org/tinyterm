//go:build calc
// +build calc

package tinyterm

import (
	"image/color"
	"strconv"
)

func (c Color) calcRGBA() color.RGBA {
	if c < 16 {
		v := uint8(128)
		if c > 7 {
			v = 255
			c -= 8
		}
		switch c {
		case ColorBlack:
			if v == 255 {
				return color.RGBA{0, 0, 0, 255}
			} else {
				return color.RGBA{128, 128, 128, 255}
			}
		case ColorRed:
			return color.RGBA{v, 0, 0, 255}
		case ColorGreen:
			return color.RGBA{0, v, 0, 255}
		case ColorYellow:
			return color.RGBA{v, v, 0, 255}
		case ColorBlue:
			return color.RGBA{0, 0, v, 255}
		case ColorMagenta:
			return color.RGBA{v, 0, v, 255}
		case ColorCyan:
			return color.RGBA{0, v, v, 255}
		case ColorWhite:
			return color.RGBA{v, v, v, 255}
		}
	}
	if c < 232 {
		// calculate rgb from 6x6x6
		var r, g, b, v uint8
		v = uint8(c - 16)
		r = v / 36
		g = (v % 36) / 6
		b = v % 6
		f := func(v uint8) uint8 {
			if v == 0 {
				return 0
			}
			return 95 + (v * 40)
		}
		return color.RGBA{f(r), f(g), f(b), 255}
	}
	if c <= 255 {
		// calculate grayscale
		v := uint8((23-(255-c))*10 + 8)
		return color.RGBA{v, v, v, 255}
	}
	panic("invalid color: " + strconv.Itoa(int(c)))
}
