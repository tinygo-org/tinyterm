package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/fonts/proggy"
)

var (
	display = st7789.New(machine.SPI1,
		machine.TFT_RESET,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_LITE)

	terminal = tinyterm.NewTerminal(&display)

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	font = &proggy.TinySZ8pt7b
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 8000000,
		SCK:       machine.TFT_SCK,
		SDO:       machine.TFT_SDO,
		SDI:       machine.TFT_SDO,
		Mode:      0,
	})
	display.Configure(st7789.Config{
		Rotation:   st7789.ROTATION_90,
		RowOffset:  80,
		FrameRate:  st7789.FRAMERATE_111,
		VSyncLines: st7789.MAX_VSYNC_SCANLINES,
	})

	width, height := display.Size()
	_, _ = width, height

	display.FillScreen(black)

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
