// This example opens a TCP connection using a device with WiFiNINA firmware
// and sends a HTTP request to retrieve a webpage, based on the following
// Arduino example:
//
// https://github.com/arduino-libraries/WiFiNINA/blob/master/examples/WiFiWebClientRepeating/
//
package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"machine"
	"time"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/wifinina"

	"github.com/conejoninja/tinyfont/proggy"
	"github.com/valyala/fastjson"

	"github.com/bgould/http/client"
	"github.com/bgould/tinyterm"
)

// access point info
const ssid = ""
const pass = ""

// IP address of the server aka "hub". Replace with your own info.
const server = "numbersapi.com"

var (
	display = ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)
	console = tinyterm.NewTerminal(display)
	black   = color.RGBA{0, 0, 0, 255}
	font    = &proggy.TinySZ8pt7b

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor = &wifinina.Device{
		SPI:   machine.NINA_SPI,
		CS:    machine.NINA_CS,
		ACK:   machine.NINA_ACK,
		GPIO0: machine.NINA_GPIO0,
		RESET: machine.NINA_RESETN,
	}

	buf  = make([]byte, 256)
	last time.Time

	buffer = bytes.NewBuffer(make([]byte, 1024))
	conn   net.Conn

	parser = &fastjson.Parser{}
)

func main() {

	machine.TFT_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})

	display.Configure(ili9341.Config{})
	display.FillScreen(black)
	machine.TFT_BACKLIGHT.High()

	console.Configure(&tinyterm.Config{
		Font:       font,
		FontHeight: 10,
		FontOffset: 6,
	})

	// Configure SPI for 8Mhz, Mode 0, MSB First
	machine.NINA_SPI.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		MOSI:      machine.NINA_MOSI,
		MISO:      machine.NINA_MISO,
		SCK:       machine.NINA_SCK,
	})

	adaptor.Configure()

	connectToAP()

	for {
		if time.Now().Sub(last).Milliseconds() >= 10000 {
			makeHTTPRequest()
		}
	}

}

var i = -1

func makeHTTPRequest() {

	i++

	var err error
	if conn != nil {
		conn.Close()
	}

	// make TCP connection
	ip := net.ParseIP(server)
	raddr := &net.TCPAddr{IP: ip, Port: 80}
	laddr := &net.TCPAddr{Port: 8080}

	message("\n---------------\nDialing TCP connection\n")
	conn, err = net.DialTCP("tcp", laddr, raddr)
	for ; err != nil; conn, err = net.DialTCP("tcp", laddr, raddr) {
		message("connection failed: " + err.Error())
		time.Sleep(5 * time.Second)
	}
	//defer conn.Close()

	message("Connected!")

	cl := client.NewClient(conn)

	cl.WriteRequest(&client.Request{
		Method:  "GET",
		Path:    fmt.Sprintf("/%d", i),
		Query:   nil,
		Version: client.HTTP_1_1,
		Headers: []client.Header{
			{"Host", server},
			{"User-Agent", "TinyGo"},
			{"Content-Type", "application/json"},
			{"Connection", "close"},
		},
		Body: nil,
	})

	last = time.Now()
	for time.Now().Sub(last).Milliseconds() < 1500 {
	}

	rsp, err := cl.ReadResponse()
	if err != nil {
		fmt.Fprintf(console, "error response: %v\n", err)
		return
	}

	buffer.Reset()
	contentLength := int(rsp.ContentLength())
	message(fmt.Sprintf("HTTP response status: %d\n", rsp.Code))
	message(fmt.Sprintf("Content length: %d\n", contentLength))

	if _, err := io.CopyBuffer(buffer, rsp.Body, buf); err != nil {
		message("error reading response body: " + err.Error())
		return
	}

	message(string(buffer.Bytes()))

	v, err := parser.ParseBytes(buffer.Bytes())
	if err != nil {
		message("error parsing JSON: " + err.Error() + "\n")
		return
	}

	message(v.Get("text").String() + "\n")
}

// connect to access point
func connectToAP() {
	time.Sleep(2 * time.Second)
	message("Connecting to " + ssid)
	adaptor.SetPassphrase(ssid, pass)
	for st, _ := adaptor.GetConnectionStatus(); st != wifinina.StatusConnected; {
		message("Connection status: " + st.String())
		time.Sleep(1 * time.Second)
		st, _ = adaptor.GetConnectionStatus()
	}
	message("Connected.")
	time.Sleep(2 * time.Second)
	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		message(err.Error())
		time.Sleep(1 * time.Second)
	}
	message(ip.String())
}

func message(msg string) {
	fmt.Fprintln(console, msg)
}
