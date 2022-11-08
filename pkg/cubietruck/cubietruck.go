package cubietruck

import (
	"github.com/jesusrj/cubietruck/pkg/cubietruck/led"
)

type Cubietruck interface {
	LedStatus(l led.Led) (led.Status, error)
	LedOn(l led.Led) error
	LedOff(l led.Led) error
}

type cubietruck struct{}

func (c *cubietruck) LedStatus(l led.Led) (led.Status, error) {
	return led.ReadStatus(l)
}

func (c *cubietruck) LedOn(l led.Led) error {
	return led.SetStatus(l, led.On)
}

func (c *cubietruck) LedOff(l led.Led) error {
	return led.SetStatus(l, led.Off)
}

func New() Cubietruck {
	return &cubietruck{}
}
