package main

import (
	"fmt"
	"image/color"
	"time"

	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/examples/initdisplay"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	font = &proggy.TinySZ8pt7b
)

func main() {
	display := initdisplay.InitDisplay()
	terminal := tinyterm.NewTerminal(display)

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
