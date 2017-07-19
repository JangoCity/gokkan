package data

import (
	"encoding/binary"
	"fmt"
	"github.com/derekparker/delve/pkg/dwarf/frame"
	"github.com/michey/gokkan/messages"
)

type IDE int32

const (
	MIN_STD_ID = 0x00
	MAX_STD_ID = 0x7FF
)

//CAN Identifier Type
const (
	CAN_IDE_STD IDE = 0
	CAN_IDE_EXT IDE = 4
)

type RTR int32

//CAN Remote Transmission Request
const (
	CAN_RTR_DATA   RTR = 0
	CAN_RTR_REMOTE RTR = 2
)

type CANFrame struct {
	StdId      uint32
	ExtendedId uint32
	IDE        IDE
	RTR        RTR
	DLC        uint32
	Data       []byte
}

func SimpleCANFrameConstruct(StdId uint32, DLC uint32, Data []byte) CANFrame {
	return CANFrame{StdId: StdId, ExtendedId: 0, IDE: CAN_IDE_STD, RTR: CAN_RTR_DATA, DLC: DLC, Data: Data}
}

func (frame *CANFrame) String() string {
	id := frame.StdId

	if frame.ExtendedId != 0 {
		id = frame.ExtendedId
	}

	d := frame.Data
	data := fmt.Sprintf("data:%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v ", d[0], d[1], d[2], d[3], d[4], d[5], d[6], d[7])

	return fmt.Sprintf("id:%#v; dlc:%d; ", id, frame.DLC) + data
}

func (frame *CANFrame) FromDevice(device *messages.FromDevice) {
	frame.StdId = device.Frame.Id
	frame.ExtendedId = device.Frame.Eid
	frame.DLC = device.Frame.Dlc
	frame.RTR = getFrameRTR(device.Frame.Rtr)
	frame.IDE = getFrameIDE(device.Frame.Ide)
	binary.LittleEndian.PutUint64(frame.Data, device.Frame.Data)
}

func (frame *CANFrame) ToDevice() messages.ToDevice {
	toDevice := messages.ToDevice{}
	toDevice.Type = messages.ToDevice_SEND_FRAME
	toDevice.Frame.Id = frame.StdId
	toDevice.Frame.Eid = frame.ExtendedId
	toDevice.Frame.Dlc = frame.DLC
	toDevice.Frame.Ide = frame.getMsgIDE()
	toDevice.Frame.Rtr = frame.getMsgRTR()
	return toDevice
}

func (frame *CANFrame) getData() uint64 {
	return binary.LittleEndian.Uint64(frame.Data)
}

func (frame *CANFrame) getMsgIDE() messages.Frame_IDE {
	if frame.IDE == CAN_IDE_STD {
		return messages.Frame_STD
	} else {
		return messages.Frame_EXT
	}
}

func (frame *CANFrame) getMsgRTR() messages.Frame_RTR {
	if frame.RTR == CAN_RTR_DATA {
		return messages.Frame_DATA
	} else {
		return messages.Frame_REMOTE
	}
}

func getFrameRTR(rtr messages.Frame_RTR) RTR {
	if rtr == messages.Frame_DATA {
		return CAN_RTR_DATA
	} else {
		return CAN_RTR_REMOTE
	}
}

func getFrameIDE(ide messages.Frame_IDE) IDE {
	if ide == messages.Frame_EXT {
		return CAN_IDE_EXT
	} else {
		return CAN_IDE_STD
	}
}
