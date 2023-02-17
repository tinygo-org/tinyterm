# tinyterm - TinyGo Terminal Emulator

[![Build](https://github.com/tinygo-org/tinyterm/actions/workflows/build.yml/badge.svg?branch=dev)](https://github.com/tinygo-org/tinyterm/actions/workflows/build.yml)

A minimal terminal for TinyGo devices supporting 256-color ANSI escape codes.

![examples/colors/main.go running on PyPortal](/examples/colors/pyportal_256color.png?raw=true)

## How to compile examples

```
tinygo flash -target pyportal -ldflags="-X main.ssid=MYSSID -X main.pass=MYPASS" ./examples/httpclient
```

