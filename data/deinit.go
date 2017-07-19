package data

import "github.com/michey/gokkan/messages"

type DeInit struct{}

func GetDeInit() Init {
	return Init{}
}

func (init *DeInit) ToDevice() messages.ToDevice {
	msg := messages.ToDevice{}
	msg.Type = messages.ToDevice_DEINIT
	return msg
}
