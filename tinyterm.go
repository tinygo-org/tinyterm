package tinyterm

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

func NewTerminal(display Displayer) *Terminal {
	return &Terminal{
		display: display,
	}
}

type Displayer interface {
	drivers.Displayer
	FillRectangle(x, y, width, height int16, c color.RGBA) error
	SetScroll(line int16)
}

type Terminal struct {
	display Displayer
	width   int16
	height  int16
	scroll  int16

	//buffer []byte // buffer for the next that has been written to the terminal
	rows int16 // number of rows in the text buffer
	cols int16 // number of columns in the text buffer
	next int16 // index in the buffer at which next char will be put

	state   state
	command byte
	params  *bytes.Buffer
	attrs   sgrAttrs

	font       *tinyfont.Font
	fontWidth  int16
	fontHeight int16
	fontOffset int16
}

type Config struct {
	ScreenBounds image.Rectangle
	Font         *tinyfont.Font
	FontHeight   int16
	FontOffset   int16
}

func (t *Terminal) Configure(config *Config) {

	t.state = stateInput
	t.params = bytes.NewBuffer(make([]byte, 32))

	t.attrs.reset()

	_, charWidth := tinyfont.LineWidth(config.Font, "0")

	t.font = config.Font
	t.fontWidth = int16(charWidth)
	t.fontHeight = config.FontHeight
	t.fontOffset = config.FontOffset

	t.width, t.height = t.display.Size()
	t.rows = t.height / t.fontHeight
	t.cols = t.width / t.fontWidth

	//t.buffer = make([]byte, t.rows*t.cols)
	t.scroll = t.fontHeight
	t.lf()
}

func (t *Terminal) Write(buf []byte) (int, error) {
	for _, b := range buf {
		t.putchar(b)
	}
	return len(buf), nil
}

func (t *Terminal) WriteByte(b byte) error {
	t.putchar(b)
	return nil
}

func (t *Terminal) Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(t, format, args...)
}

func (t *Terminal) Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(t, args...)
}

type state uint8

const (
	stateInput state = iota
	stateEscape
	stateCSI
)

func (t *Terminal) putchar(b byte) {
	switch t.state {
	case stateInput:
		switch b {
		case 0x1b:
			t.state = stateEscape
			return
		case '\r':
			t.cr()
			return
		case '\n':
			t.lf()
			return
		default:
			t.drawchar(b)
			return
		}

	case stateEscape:
		switch b {
		case 'N':
			// SS2: Single Shift Two
			t.state = stateInput
		case 'O':
			// SS3: Single Shift Three
			t.state = stateInput
		case 'P':
			// DCS: Device Control String
			t.command = b
		case '[':
			// CSI: Control Sequence Introducer
			t.params.Reset()
			t.state = stateCSI
		case '\\':
			// ST: String Terminator
			t.state = stateInput
		case ']':
			// OSC: Operating System Command
			t.command = b
		case 'X':
			// SOS: Start of String
			t.command = b
		case '^':
			// PM: Privacy Message
			t.command = b
		case '_':
			// APC: Application Program Command
			t.command = b
		case 'c':
			// RIS: Reset to Initial State
			// TODO: need to implement
			t.state = stateInput
		}
	case stateCSI:
		switch {
		case b >= 0x20 && b <= 0x2F:
			// intermediate bytes
			t.params.WriteByte(b)
		case b >= 0x30 && b <= 0x3F:
			// parameter bytes
			t.params.WriteByte(b)
		default:
			// final bytes
			switch b {
			case 'A':
				// CUU: Cursor Up
			case 'B':
				// CUD: Cursor Down
			case 'C':
				// CUF: Cursor Forward
			case 'D':
				// CUB: Cursor Back
			case 'E':
				// CNL: Cursor Next Line
			case 'F':
				// CPL: Cursor Previous Line
			case 'G':
				// CHA: Cursor Horizontal Absolute
			case 'H':
				// CUP: Cursor Position
			case 'J':
				// ED: Erase in Display
			case 'K':
				// EL: Erase in Line
			case 'S':
				// SU: Scroll Up
			case 'T':
				// SD: Scroll Down
			case 'f':
				// HVP: Horizontal Vertical Position
			case 'm':
				// SGR: Select Graphic Rendition
				t.selectGraphicRendition()
			case 'i':
				// AUX Port
			case 'n':
				// DSR: Device Status Report
			default:
				// undefined behavior; just reset the sequence to input mode
			}
			t.state = stateInput
		}
	}
}

func (t *Terminal) selectGraphicRendition() {
	params := strings.Split(t.params.String(), ";")
	attr, err := strconv.Atoi(params[0])
	if err != nil {
		println("error converting SGR param ID: " + err.Error())
	}
	switch attr {
	case SGRReset:
		t.attrs.reset()
	case SGRBold:
		t.attrs.attrs |= byte(attr)
	case SGRFgBlack:
		fallthrough
	case SGRFgRed:
		fallthrough
	case SGRFgGreen:
		fallthrough
	case SGRFgYellow:
		fallthrough
	case SGRFgBlue:
		fallthrough
	case SGRFgMagenta:
		fallthrough
	case SGRFgCyan:
		fallthrough
	case SGRFgWhite:
		t.attrs.setFG(Color(attr % 10))
	case SGRSetFgColor:
		c, err := strconv.Atoi(params[2])
		if err != nil {
			println("error converting color: " + err.Error())
		}
		t.attrs.setFG(Color(c))
	case SGRDefaultFgColor:
		t.attrs.setFG(ColorWhite)
	case SGRBgBlack:
		fallthrough
	case SGRBgRed:
		fallthrough
	case SGRBgGreen:
		fallthrough
	case SGRBgYellow:
		fallthrough
	case SGRBgBlue:
		fallthrough
	case SGRBgMagenta:
		fallthrough
	case SGRBgCyan:
		fallthrough
	case SGRBgWhite:
		t.attrs.setBG(Color(attr % 10))
	case SGRSetBgColor:
		c, err := strconv.Atoi(params[2])
		if err != nil {
			println("error converting color: " + err.Error())
		}
		t.attrs.setBG(Color(c))
	case SGRDefaultBgColor:
		t.attrs.setBG(ColorBlack)
	}
}

func (t *Terminal) drawchar(b byte) {
	if t.next == t.cols {
		t.lf()
	}
	x := t.next * t.fontWidth
	y := t.scroll + t.fontOffset
	t.display.FillRectangle(x, t.scroll, t.fontWidth, t.fontHeight, t.attrs.bgcol)
	tinyfont.DrawChar(t.display, t.font, x, y, rune(b), t.attrs.fgcol)
	t.next += 1
}

func (t *Terminal) cr() {
	//t.next = 0
	//t.display.FillRectangle(0, t.scroll, t.width, t.fontHeight, black)
}

func (t *Terminal) lf() {
	t.next = 0
	t.scroll = (t.scroll + t.fontHeight) % (t.rows * t.fontHeight)
	t.display.SetScroll((t.scroll + t.fontHeight) % t.height)
	t.display.FillRectangle(0, t.scroll, t.width, t.fontHeight, t.attrs.bgcol)
	t.display.Display()
}
