package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/fonts/proggy"
)

var (
	terminal = tinyterm.NewTerminal(display)

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	font = &proggy.TinySZ8pt7b
)

func main() {

	backlight.Configure(machine.PinConfig{machine.PinOutput})

	width, height := display.Size()
	_, _ = width, height

	display.FillScreen(black)
	backlight.High()

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
