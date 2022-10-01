package main

import (
	"fmt"
	"image/color"
	"machine"
	"strings"
	"time"

	"tinygo.org/x/drivers/ili9341"

	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
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

	font = &proggy.TinySZ8pt7b
)

func main() {

	time.Sleep(time.Second)

	machine.TFT_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})

	display.Configure(ili9341.Config{})
	width, height := display.Size()
	_, _ = width, height

	display.FillScreen(color.RGBA{0, 0, 0, 255})
	machine.TFT_BACKLIGHT.High()

	terminal.Configure(&tinyterm.Config{
		Font:       font,
		FontHeight: 10,
		FontOffset: 6,
	})

	for n := 0; ; n++ {
		terminal.Write([]byte("   " + strings.Repeat("_", 36) + "\n"))
		fmt.Fprintf(terminal, "%02x|", 0)
		for n := 0; n < 16; n++ {
			fmt.Fprintf(terminal, "\x1b[48;5;%dm \x1b[0m", n)
		}
		terminal.Write([]byte(strings.Repeat(" ", 20) + "|\n"))
		for n := 0; n < 6; n++ {
			i := n*36 + 16
			fmt.Fprintf(terminal, "%02x|", i)
			for j := 0; j < 36; j++ {
				v := i + j
				fmt.Fprintf(terminal, "\x1b[48;5;%dm \x1b[0m", v)
			}
			terminal.WriteByte('|')
			terminal.WriteByte('\n')
		}
		fmt.Fprintf(terminal, "%02x|", 232)
		for n := 232; n <= 255; n++ {
			fmt.Fprintf(terminal, "\x1b[48;5;%dm \x1b[0m", n)
		}
		terminal.Write([]byte(strings.Repeat(" ", 12) + "|\n"))
		terminal.Write([]byte("   " + strings.Repeat("\xaf", 36) + "\n"))
		time.Sleep(5 * time.Second)
	}

}
