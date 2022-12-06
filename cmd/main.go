package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jesusrj/cubietruck/pkg/cubietruck"
)

const (
	// ANSI Escape codes
	CLEAR_SCREEN = "\033[H\033[2J"
	CLEAR_LINE   = "\033[2K\r"

	// UNICODE (https://unicode.org/emoji/charts/emoji-list.html)
	// https://dev.to/matthewdale/sending-in-go-46bf
	ICO_THERMOMETER = '\U0001F321'
	ICO_HOME        = '\U0001F3E0'
)

func main() {
	fmt.Println("---- Cubietruck Manager ----")
	var ct = cubietruck.New()

	go printCPUTemp(ct)
	go printIPAddress()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
}

func printCPUTemp(ct cubietruck.Cubietruck) {
	if t, err := ct.CPUTemp(); err != nil {
		return
	} else {
		fmt.Printf("%s %c  CPU Temp: %dC", CLEAR_LINE, ICO_THERMOMETER, t/1000)
	}

	time.Sleep(1 * time.Second)
	go printCPUTemp(ct)
}

func printIPAddress() {
	addr, err := getIPAddress()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %c  IP: %s\n", CLEAR_LINE, ICO_HOME, addr.String())
}

func getIPAddress() (net.Addr, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	conn.Close()

	return conn.LocalAddr(), nil
}
