package main

import (
	"fmt"
	"image/color"
	"time"

	"machine"

	"tinygo.org/x/drivers/st7735"

	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
)

var (
	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)

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
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	width, height := display.Size()
	_, _ = width, height

	display.FillScreen(black)

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})
	for {
		time.Sleep(50 * time.Millisecond)
		fmt.Fprintf(terminal, "\ntime: %d", time.Now().UnixNano())
	}
}
