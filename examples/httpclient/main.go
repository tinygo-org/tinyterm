// This example requires a network adaptor. It gets an URL using http.Get() and
// displays the output on the terminal screen.
//
// Note: It may be necessary to increase the stack size when using "net/http".
// Use the -stack-size=4KB command line option.

package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"tinygo.org/x/drivers/netlink"
	"tinygo.org/x/drivers/netlink/probe"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/examples/initdisplay"
)

var (
	ssid string
	pass string
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	font = &proggy.TinySZ8pt7b
)

func main() {
	display := initdisplay.InitDisplay()
	terminal := tinyterm.NewTerminal(display)

	terminal.Configure(&tinyterm.Config{
		Font:       font,
		FontHeight: 10,
		FontOffset: 6,
	})

	fmt.Fprintf(terminal, "Connecting to %s...\r\n", ssid)

	link, _ := probe.Probe()

	err := link.NetConnect(&netlink.ConnectParams{
		Ssid:       ssid,
		Passphrase: pass,
	})
	if err != nil {
		log.Fatal(err)
	}

	name := "John Doe"
	occupation := "gardener"

	params := "name=" + url.QueryEscape(name) + "&" +
		"occupation=" + url.QueryEscape(occupation)

	path := fmt.Sprintf("https://httpbin.org/get?%s", params)

	cnt := 0
	for {
		fmt.Fprintf(terminal, "Getting %s\r\n\r\n", path)
		resp, err := http.Get(path)
		if err != nil {
			fmt.Fprintf(terminal, "%s\r\n", err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Fprintf(terminal, "%s %s\r\n", resp.Proto, resp.Status)
		for k, v := range resp.Header {
			fmt.Fprintf(terminal, "%s: %s\r\n", k, strings.Join(v, " "))
		}
		fmt.Fprintf(terminal, "\r\n")

		body, err := io.ReadAll(resp.Body)
		fmt.Fprintf(terminal, string(body))
		resp.Body.Close()

		cnt++
		fmt.Fprintf(terminal, "-------- %d --------\r\n", cnt)
		time.Sleep(10 * time.Second)
	}
}
