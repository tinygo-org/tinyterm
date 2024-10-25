package main

import (
	"fmt"
	"time"

	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/examples/initdisplay"
)

var (
	font = &proggy.TinySZ8pt7b
)

func main() {
	display := initdisplay.InitDisplay()
	terminal := tinyterm.NewTerminal(display)

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: initdisplay.NeedsSoftwareScroll(),
	})
	for {
		time.Sleep(time.Second)

		fmt.Fprintf(terminal, "\ntime: %d", time.Now().UnixNano())
		display.Display()
	}
}
