# tinyterm - TinyGo Terminal Emulator

[![Build](https://github.com/tinygo-org/tinyterm/actions/workflows/build.yml/badge.svg?branch=dev)](https://github.com/tinygo-org/tinyterm/actions/workflows/build.yml)

A minimal terminal for TinyGo displays. Supporting 256-color ANSI escape codes, as well as monochrome displays such as e-ink or OLED.

![examples/colors/main.go running on PyPortal](/examples/colors/pyportal_256color.png?raw=true)

## How to use it

```go
package main

import (
	"fmt"
	"time"

	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/displays"
)

var (
	font = &proggy.TinySZ8pt7b
)

func main() {
	display := displays.Init()
	terminal := tinyterm.NewTerminal(display)

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: displays.NeedsSoftwareScroll(),
	})
	for {
		time.Sleep(time.Second)

		fmt.Fprintf(terminal, "\ntime: %d", time.Now().UnixNano())
		terminal.Display()
	}
}
```

## How to compile examples

Most of the examples will work with any of the following hardware:

- Adafruit Clue (https://www.adafruit.com/clue)
- Badger2040 & Badger2040-W (https://shop.pimoroni.com/products/badger-2040)
- Gopher Badge (https://gopherbadge.com/)
- PyBadge (https://www.adafruit.com/product/4200)
- PyPortal (https://www.adafruit.com/product/4116)
- WioTerminal (https://wiki.seeedstudio.com/Wio-Terminal-Getting-Started/)

### basic

Displays basic text.

```
tinygo flash -target pyportal ./examples/basic
```

### colors

Displays ANSI colors.

```
tinygo flash -target pyportal ./examples/colors
```

### httpclient

Connects to an http server and displays the results. Runs on PyPortal and WioTerminal only, since it requires a connected WiFi coprocessor.

```
tinygo flash -target pyportal -ldflags="-X main.ssid=MYSSID -X main.pass=MYPASS" ./examples/httpclient
```

