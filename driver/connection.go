package driver

import (
	"github.com/michey/gokkan/messages"
)

type CANConnection interface {
	Send(frame messages.ToDevice)
	GetChan() chan messages.FromDevice
}
