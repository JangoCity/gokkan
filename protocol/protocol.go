package protocol

import (
	"github.com/michey/gokkan/data"
	"github.com/michey/gokkan/messages"
)

type IProtocol interface {
	Encode(frame *messages.ToDevice) []byte
	Decode(bytes <-chan byte, frames chan<- messages.FromDevice)
	Stop()
}
