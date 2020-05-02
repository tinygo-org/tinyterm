package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"github.com/bgould/tinyterm"
	"github.com/bgould/tinyterm/fonts/proggy"
	"tinygo.org/x/drivers/ili9341"
)

var (
	display = ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)

	terminal = tinyterm.NewTerminal(display)

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	font = &proggy.TinySZ8pt7b
)

func main() {

	machine.TFT_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})

	display.Configure(ili9341.Config{})
	width, height := display.Size()
	_, _ = width, height

	display.FillScreen(black)
	machine.TFT_BACKLIGHT.High()

	terminal.Configure(&tinyterm.Config{
		Font:       font,
		FontHeight: 10,
		FontOffset: 6,
	})
	for {
		time.Sleep(50 * time.Millisecond)
		fmt.Fprintf(terminal, "\ntime: %d", time.Now().UnixNano())
	}

}
