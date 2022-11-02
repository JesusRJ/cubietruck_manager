package led

import "github.com/jesusrj/cubietruck/pkg/device"

// Led represent a device led brightness
type Led interface {
	device.Input
	device.Output
}

type led struct {
	fd string
}

func New(fd string) Led {
	return &led{
		fd: fd,
	}
}

func (l *led) Read() ([]byte, error) {
	return nil, nil
}

func (l *led) Write(b []byte) (int, error) {
	return 0, nil
}
