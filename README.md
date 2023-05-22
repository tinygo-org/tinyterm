# tinyterm - TinyGo Terminal Emulator

[![Build](https://github.com/tinygo-org/tinyterm/actions/workflows/build.yml/badge.svg?branch=dev)](https://github.com/tinygo-org/tinyterm/actions/workflows/build.yml)

A minimal terminal for TinyGo devices supporting 256-color ANSI escape codes.

![examples/colors/main.go running on PyPortal](/examples/colors/pyportal_256color.png?raw=true)

## How to compile examples

Most of the examples will work with any of the following hardware:

- Adafruit Clue (https://www.adafruit.com/clue)
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

