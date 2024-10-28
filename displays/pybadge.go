//go:build pybadge

package displays

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyterm"
)

func Init() tinyterm.Displayer {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	return &display
}

func NeedsSoftwareScroll() bool {
	return false
}
