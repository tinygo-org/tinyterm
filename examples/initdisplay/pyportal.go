//go:build pyportal

package initdisplay

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinyterm"
)

func InitDisplay() tinyterm.Displayer {
	display := ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)

	// configure backlight
	backlight := machine.TFT_BACKLIGHT
	backlight.Configure(machine.PinConfig{machine.PinOutput})

	// configure display
	display.Configure(ili9341.Config{})
	display.SetRotation(ili9341.Rotation0)
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	backlight.High()

	return display
}
