package protocol

import (
	"github.com/michey/gokkan/data"
)

type IProtocol interface {
	Encode(frame *data.CANFrame) []byte
	Decode(bytes []byte) (err error, frame *data.CANFrame)
}
