package temp

import "github.com/jesusrj/cubietruck/pkg/device"

// Thermal represent a device temperature
type Thermal interface {
	device.Output
}

type temp struct {
	fd string
}

func New(fd string) Thermal {
	return &temp{
		fd: fd,
	}
}

func (t *temp) Write(b []byte) (int, error) {
	return 0, nil
}
