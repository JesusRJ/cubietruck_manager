package device

// Input is a interface that represent a input device
type Input interface {
	Read() ([]byte, error)
}

// Output is a interface that represetn a output device
type Output interface {
	Write(b []byte) (int, error)
}
