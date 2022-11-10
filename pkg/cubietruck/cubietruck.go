package cubietruck

import (
	"time"

	"github.com/jesusrj/cubietruck/pkg/cubietruck/led"
	"github.com/jesusrj/cubietruck/pkg/cubietruck/thermal"
)

type Cubietruck interface {
	LedStatus(l led.Led) (led.Status, error)
	LedOn(l led.Led) error
	LedOff(l led.Led) error
	LedBlink(l led.Led, d time.Duration) error
	CPUTemp() (uint, error)
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

func (c *cubietruck) LedBlink(l led.Led, d time.Duration) error {
	err := c.LedOn(l)
	if err != nil {
		return err
	}

	time.Sleep(d)

	err = c.LedOff(l)
	if err != nil {
		return err
	}
	return nil
}

func (c *cubietruck) CPUTemp() (uint, error) {
	c.LedOn(led.Orange)
	x, err := thermal.CPUTemp()
	c.LedOff(led.Orange)
	return x, err
}

func New() Cubietruck {
	return &cubietruck{}
}
