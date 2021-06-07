// +build wioterminal

package main

import (
	"machine"

	"tinygo.org/x/drivers/ili9341"
)

var (
	display = ili9341.NewSPI(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)

	backlight = machine.LCD_BACKLIGHT
)

func init() {
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 40000000,
	})
	display.Configure(ili9341.Config{})
}
