package data

import "github.com/michey/gokkan/messages"

type IData interface {
	ToDevice() messages.ToDevice
}

type PopulatableData interface {
	FromDevice(from *messages.FromDevice)
}
