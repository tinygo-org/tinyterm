//go:build clue_alpha

package displays

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyterm"
)

func Init() tinyterm.Displayer {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 8000000,
		SCK:       machine.TFT_SCK,
		SDO:       machine.TFT_SDO,
		SDI:       machine.TFT_SDO,
		Mode:      0,
	})

	display := st7789.New(machine.SPI1,
		machine.TFT_RESET,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_LITE)

	display.Configure(st7789.Config{
		Rotation:   st7789.ROTATION_180,
		Height:     320,
		FrameRate:  st7789.FRAMERATE_111,
		VSyncLines: st7789.MAX_VSYNC_SCANLINES,
	})

	display.FillScreen(color.RGBA{0, 0, 0, 255})

	return &display
}

func NeedsSoftwareScroll() bool {
	return false
}
