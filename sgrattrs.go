package tinyterm

import (
	"image/color"
)

const (
	SGRReset = iota
	SGRBold
	SGRFaint
	SGRItalic
	SGRUnderline
	SGRSlowBlink
	SGRRapidBlink
	SGRReverseVideo
	SGRConceal
	SGRCrossedOut
	SGRPrimaryFont
	SGRAltFont1
	SGRAltFont2
	SGRAltFont3
	SGRAltFont4
	SGRAltFont5
	SGRAltFont6
	SGRAltFont7
	SGRAltFont8
	SGRAltFont9
	SGRFraktur
	SGRDoubleUnderline
	SGRNormal
	SGRNotItalicNotFraktur
	SGRUnderlineOff
	SGRBlinkOff
	SGRInverseOff
	SGRReveal
	SGRNotCrossedOut
	_
	SGRFgBlack
	SGRFgRed
	SGRFgGreen
	SGRFgYellow
	SGRFgBlue
	SGRFgMagenta
	SGRFgCyan
	SGRFgWhite
	SGRSetFgColor
	SGRDefaultFgColor
	SGRBgBlack
	SGRBgRed
	SGRBgGreen
	SGRBgYellow
	SGRBgBlue
	SGRBgMagenta
	SGRBgCyan
	SGRBgWhite
	SGRSetBgColor
	SGRDefaultBgColor
	SGRFramed
	SGREncircled
	SGROverlined
	SGRIdeogramUnderline
	SGRIdeogramDblUnderline
	SGRIdeogramOverline
	SGRIdeogramStress
	SGRIdeogramAttrOff
)

type Color uint8

const (
	ColorBlack Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	ColorBrBlack
	ColorBrRed
	ColorBrGreen
	ColorBrYellow
	ColorBrBlue
	ColorBrMagenta
	ColorBrCyan
	ColorBrWhite
)

type sgrAttrs struct {
	fg    Color
	bg    Color
	attrs uint8

	fgcol color.RGBA
	bgcol color.RGBA
}

func (a *sgrAttrs) setFG(c Color) {
	a.fg = c
	a.fgcol = a.calcColor(c).calcRGBA()
}

func (a *sgrAttrs) setBG(c Color) {
	a.bg = c
	a.bgcol = a.calcColor(c).calcRGBA()
}

func (a *sgrAttrs) calcColor(c Color) Color {
	if c < 8 && a.attrs&SGRBold > 0 {
		c += 8
	}
	return c
}

func (a *sgrAttrs) reset() {
	a.attrs = 0
	a.setFG(ColorBrWhite)
	a.setBG(ColorBlack)
}
