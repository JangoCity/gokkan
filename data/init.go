package data

import "github.com/michey/gokkan/messages"

type Init struct{}

func GetInit() Init {
	return Init{}
}

func (init *Init) ToDevice() messages.ToDevice {
	msg := messages.ToDevice{}
	msg.Type = messages.ToDevice_INIT
	return msg
}
