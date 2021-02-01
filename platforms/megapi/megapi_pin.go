package megapi

import (
	"bytes"
)

// Pin represents a pin
type Pin struct {
	megaPi    *Adaptor
	pinNumber byte
}

// NewPin creates a new PinDriver at the given pin number
func NewPin(megaPi *Adaptor, pinNumber byte) *Pin {
	return &Pin{
		megaPi:    megaPi,
		pinNumber: pinNumber,
	}
}

// On turns power on the pin
func (pin *Pin) On() {
	pin.digitalWrite(0x1)
}

// Off turns power off the pin
func (pin *Pin) Off() {
	pin.digitalWrite(0x0)
}

func (pin *Pin) digitalWrite(val byte) (err error) {
	bufOut := new(bytes.Buffer)
	bufOut.Write([]byte{0xff, 0x55, 0x5, 0x0, 0x2, 0x1e, pin.pinNumber, val})
	pin.megaPi.writeBytesChannel <- bufOut.Bytes()
	return nil
}
