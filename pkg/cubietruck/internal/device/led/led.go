package led

import (
	"os"

	"github.com/jesusrj/cubietruck/pkg/cubietruck/internal/device"
)

// Led represent a device led brightness
type Led interface {
	device.Reader
	device.Writer
}

type led struct {
	fd string
}

func (l *led) Read() ([]byte, error) {
	return os.ReadFile(l.fd)
}

func (l *led) Write(b []byte) error {
	return os.WriteFile(l.fd, b, 0644)
}

func New(fd string) Led {
	return &led{
		fd: fd,
	}
}
