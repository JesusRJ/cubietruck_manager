package thermal

import (
	"github.com/jesusrj/cubietruck/pkg/cubietruck/internal/device"
)

// Thermal represent a device temperature
type Thermal interface {
	device.Reader
}

type temp struct {
	fd string
}

func (t *temp) Read() ([]byte, error) {
	return nil, nil
}

func New(fd string) Thermal {
	return &temp{
		fd: fd,
	}
}
