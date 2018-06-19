package uds

import "github.com/michey/gokkan/protocol"

type ToFrame interface {
	ToFrame(id uint32) protocol.CANFrame
}

type SingleFrame struct {
	SID    SID
	Data   []byte
	Length uint8
}

func (sf *SingleFrame) ToFrame(id uint32) *protocol.CANFrame {
	//well, we can cheat. just use length. but by spec - [0,0,0,0, ..4 bit of length..]
	b1 := sf.Length
	b2 := sf.SID
	d := make([]byte, 8)
	copy(d[2:], sf.Data)
	d[0] = b1
	d[1] = uint8(b2)
	dlc := sf.Length + uint8(len(sf.Data))

	return &protocol.CANFrame {
		id,
		0x00,
		data.CAN_IDE_STD,
		data.CAN_RTR_DATA,
		uint32(dlc),
		d}
}
