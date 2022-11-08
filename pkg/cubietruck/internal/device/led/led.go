package led

import (
	"fmt"
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
	f, err := os.OpenFile(l.fd, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var b []byte
	n, err := f.Read(b)

	if err != nil {
		return nil, err
	}

	if n <= 0 {
		fmt.Println("N is zero")
	}

	return b, nil
}

func (l *led) Write(b []byte) (int, error) {
	f, err := os.OpenFile(l.fd, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(b)
}

func New(fd string) Led {
	return &led{
		fd: fd,
	}
}
