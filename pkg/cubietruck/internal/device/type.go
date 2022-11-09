package device

// Reader is a interface that represent a input device
type Reader interface {
	Read() ([]byte, error)
}

// Writer is a interface that represent a output device
type Writer interface {
	Write(b []byte) error
}
