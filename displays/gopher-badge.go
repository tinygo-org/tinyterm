//go:build gopher_badge

package displays

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyterm"
)

func Init() tinyterm.Displayer {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE
	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_180,
		Height:   320,
		Width:    240,
	})
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	return &display
}

func NeedsSoftwareScroll() bool {
	return false
}
