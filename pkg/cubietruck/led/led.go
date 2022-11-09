package led

import (
	"strconv"

	"github.com/jesusrj/cubietruck/pkg/cubietruck/internal/device/led"
)

const (
	led_white  = "/sys/class/leds/cubietruck:white:usr/brightness"
	led_orange = "/sys/class/leds/cubietruck:orange:usr/brightness"
	led_green  = "/sys/class/leds/cubietruck:green:usr/brightness"
	led_blue   = "/sys/class/leds/cubietruck:blue:usr/brightness"
)

type Led uint

// Leds constants
const (
	White Led = iota
	Orange
	Green
	Blue
)

type Status uint

const (
	Off Status = iota
	On
)

// Leds
var leds = map[Led]led.Led{
	White:  led.New(led_white),
	Orange: led.New(led_orange),
	Green:  led.New(led_green),
	Blue:   led.New(led_blue),
}

var Leds = [...]Led{Blue, Orange, White, Green}

// TODO: Qual seria o melhor retorno para Status?
func ReadStatus(l Led) (Status, error) {
	ledInstance := leds[l]
	b, err := ledInstance.Read()
	if err != nil {
		return Off, err
	}

	v, _ := strconv.Atoi(string(b))

	return Status(v), nil
}

func SetStatus(l Led, s Status) error {
	state := strconv.Itoa(int(s))
	ledInstance := leds[l]
	return ledInstance.Write([]byte(state))
}
