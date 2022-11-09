package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jesusrj/cubietruck/pkg/cubietruck"
	"github.com/jesusrj/cubietruck/pkg/cubietruck/led"
)

const (
	// ANSI Escape codes
	CLEAR_SCREEN = "\033[H\033[2J"
	CLEAR_LINE   = "\033[2K\r"

	// UNICODE (https://www.tamasoft.co.jp/en/general-info/unicode.html)
	// https://dev.to/matthewdale/sending-in-go-46bf
	ICO_THERMOMETER = '\U0001F321'
)

func main() {
	fmt.Println("---- Cubietruck Manager ----")
	var ct = cubietruck.New()

	go printCPUTemp(ct)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
}

func printCPUTemp(ct cubietruck.Cubietruck) {
	ct.LedOn(led.Orange)
	t, err := ct.CPUTemp()
	ct.LedOff(led.Orange)
	if err != nil {
		return
	}

	fmt.Printf("%s %c  CPU Temp: %dC", CLEAR_LINE, ICO_THERMOMETER, t/1000)
	time.Sleep(1 * time.Second)
	go printCPUTemp(ct)
}
