package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jesusrj/cubietruck/pkg/cubietruck"
)

const (
	// ANSI Escape codes
	CLEAR_SCREEN = "\033[H\033[2J"
	CLEAR_LINE   = "\033[2K\r"
)

func main() {

}

func principal() {
	leds := [...]string{cubietruck.LED_BLUE, cubietruck.LED_ORANGE, cubietruck.LED_WHITE, cubietruck.LED_GREEN}
	p := 0
	calcPos := inc
	ctemp := make(chan int)

	fmt.Println("---- Cubietruck Manager ----")

	go func() {
		go printCpuTemp(ctemp)
		for {
			blink(leds[p], time.Duration(100*time.Millisecond))

			p = calcPos(p)

			if p >= 3 {
				calcPos = dec
			}
			if p <= 0 {
				calcPos = inc
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
}

func inc(x int) int {
	return x + 1
}

func dec(x int) int {
	return x - 1
}

func printCpuTemp(c chan int) {
	go readCpuTemp(c)
	t, ok := <-c
	if !ok {
		return
	}
	fmt.Printf("%sCPU Temp: %dC", CLEAR_LINE, t/1000)
	printCpuTemp(c)
}

func readCpuTemp(c chan<- int) {
	// _, err = tempFile.Read(temp)
	temp, err := os.ReadFile(CPU_TEMP)
	if err != nil {
		log.Fatal(err)
	}

	// trim carriage return
	stemp := strings.TrimSuffix(string(temp), "\n")
	t, err := strconv.Atoi(stemp)

	if err != nil {
		log.Fatal(err)
	}
	c <- t
}

func blink(led string, d time.Duration) {
	ledFile, err := os.OpenFile(led, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer ledFile.Close()

	_, err = ledFile.WriteString("1")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(d)
	_, err = ledFile.WriteString("0")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(d)
}
