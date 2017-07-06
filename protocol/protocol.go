package protocol

import (
	"github.com/michey/gokkan/data"
)

type IProtocol interface {
	Encode(frame *data.CANFrame) []byte
	Decode(bytes <-chan byte, frames chan<- data.Response)
	Stop()
}
