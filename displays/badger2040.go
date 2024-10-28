//go:build badger2040 || badger2040_w

package displays

import (
	"machine"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyterm"
)

func Init() tinyterm.Displayer {
	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display := uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Speed:       uc8151.TURBO,
		FlickerFree: true,
		Rotation:    uc8151.ROTATION_270,
	})

	display.ClearDisplay()

	return &display
}

func NeedsSoftwareScroll() bool {
	return true
}
