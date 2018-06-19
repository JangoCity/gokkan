package protocol

import (
	"fmt"
	"bytes"
)

type CANFrame struct {
	id   uint32
	dlc  uint8
	rtr  bool
	ide  bool
	data []byte
}

func InitCanFrame(id uint32, dlc uint8, data []byte) *CANFrame {
	return &CANFrame{id, dlc, false, false, data}
}

func InitExtCanFrame(id uint32, dlc uint8, data []byte) *CANFrame {
	return &CANFrame{id, dlc, false, true, data}
}

func InitRtrCanFrame(id uint32, dlc uint8) *CANFrame {
	return &CANFrame{id, dlc, true, false, []byte{}}
}

func InitRtrExtCanFrame(id uint32, dlc uint8) *CANFrame {
	return &CANFrame{id, dlc, true, true, []byte{}}
}

func (c *CANFrame) Decode() []byte {
	var strBuilder bytes.Buffer
	if (!c.ide && !c.rtr) {
		strBuilder.WriteString("t")
		strBuilder.WriteString(fmt.Sprintf("%03x", c.id))
		strBuilder.WriteString(fmt.Sprintf("%d", c.dlc))
		for i := uint8(0); i < c.dlc; i++ {
			strBuilder.WriteString(fmt.Sprintf("%02x", c.data[i]))
		}
	} else if (c.ide && !c.rtr) {
		strBuilder.WriteString("T")
		strBuilder.WriteString(fmt.Sprintf("%08x", c.id))
		strBuilder.WriteString(fmt.Sprintf("%d", c.dlc))
		for i := uint8(0); i < c.dlc; i++ {
			strBuilder.WriteString(fmt.Sprintf("%02x", c.data[i]))
		}
	} else if (!c.ide && c.rtr) {
		strBuilder.WriteString("r")
		strBuilder.WriteString(fmt.Sprintf("%03x", c.id))
		strBuilder.WriteString(fmt.Sprintf("%d", c.dlc))
	} else if (c.ide && c.rtr) {
		strBuilder.WriteString("R")
		strBuilder.WriteString(fmt.Sprintf("%08x", c.id))
		strBuilder.WriteString(fmt.Sprintf("%d", c.dlc))
	}
	strBuilder.WriteString("\r")
	return strBuilder.Bytes()
}

