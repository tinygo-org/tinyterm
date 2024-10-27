package main

import (
	"fmt"
	"strings"
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
		Font:       font,
		FontHeight: 10,
		FontOffset: 6,
	})

	time.Sleep(time.Second)

	for n := 0; ; n++ {
		terminal.Write([]byte("   " + strings.Repeat("_", 36) + "\n"))
		fmt.Fprintf(terminal, "%02x|", 0)
		for n := 0; n < 16; n++ {
			fmt.Fprintf(terminal, "\x1b[48;5;%dm \x1b[0m", n)
		}
		terminal.Write([]byte(strings.Repeat(" ", 20) + "|\n"))
		for n := 0; n < 6; n++ {
			i := n*36 + 16
			fmt.Fprintf(terminal, "%02x|", i)
			for j := 0; j < 36; j++ {
				v := i + j
				fmt.Fprintf(terminal, "\x1b[48;5;%dm \x1b[0m", v)
			}
			terminal.WriteByte('|')
			terminal.WriteByte('\n')
		}
		fmt.Fprintf(terminal, "%02x|", 232)
		for n := 232; n <= 255; n++ {
			fmt.Fprintf(terminal, "\x1b[48;5;%dm \x1b[0m", n)
		}
		terminal.Write([]byte(strings.Repeat(" ", 12) + "|\n"))
		terminal.Write([]byte("   " + strings.Repeat("\xaf", 36) + "\n"))

		terminal.Display()
		time.Sleep(5 * time.Second)
	}

}
