package data

import "github.com/michey/gokkan/messages"

type BaudRate struct {
	Rate messages.BaudRate_Rate
}

func (init *BaudRate) ToDevice() messages.ToDevice {
	msg := messages.ToDevice{}
	msg.Type = messages.ToDevice_SET_BAUDRATE
	msg.BaudRate.Rate = init.Rate
	return msg
}
