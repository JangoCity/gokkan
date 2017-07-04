package driver

import "github.com/michey/gokkan/data"

type CANConnection interface {
	Send(frame data.CANFrame)
	GetChan() chan data.Response
}
