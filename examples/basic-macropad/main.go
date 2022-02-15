//go:build macropad_rp2040
// +build macropad_rp2040

package main

import (
	"fmt"
	"image/color"
	"machine"

	"tinygo.org/x/drivers/sh1106"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/fonts/proggy"
)

var (
	display  = sh1106.NewSPI(machine.SPI1, machine.OLED_DC, machine.OLED_RST, machine.OLED_CS)
	terminal = tinyterm.NewTerminal(&displayer{&display})
	font     = &proggy.TinySZ8pt7b
)

func init() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 48000000,
	})
	display.Configure(sh1106.Config{
		Width:  128,
		Height: 64,
	})
	display.ClearDisplay()
	terminal.Configure(&tinyterm.Config{
		Font:       font,
		FontHeight: 8,
		FontOffset: 6,
	})
	// machine.Serial = &serialWrapper{terminal}
}

func main() {
	for i := 0; ; i++ {
		//time.Sleep(10 * time.Millisecond)
		fmt.Fprintf(terminal, "counter: %d\n", i)
	}
}

type displayer struct {
	*sh1106.Device
}

func (d *displayer) FillRectangle(x, y, w, h int16, c color.RGBA) error {
	tinydraw.FilledRectangle(d.Device, x, y, w, h, c)
	return nil
}

// type serialWrapper struct {
// 	*tinyterm.Terminal
// }

// func (sw *serialWrapper) Read(p []byte) (int, error) {
// 	return 0, io.EOF
// }

// func (sw *serialWrapper) Buffered() int {
// 	return 0
// }

// func (sw *serialWrapper) Configure(uart machine.UARTConfig) error {
// 	return nil
// }
