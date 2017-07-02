package data

import (
	"encoding/binary"
)

//CAN Identifier Type
const (
	CAN_IDE_STD = 0
	CAN_IDE_EXT = 4
)

//CAN Remote Transmission Request
const (
	CAN_RTR_DATA   = 0
	CAN_RTR_REMOTE = 2
)

type CANFrame struct {
	StdId      uint32
	ExtendedId uint32
	IDE        uint32
	RTR        uint32
	DLC        uint32
	Data       []byte
}

func SimpleCANFrameConstruct(StdId uint32, DLC uint32, Data []byte) CANFrame {
	return CANFrame{StdId: StdId, ExtendedId: 0, IDE: CAN_IDE_STD, RTR: CAN_RTR_DATA, DLC: DLC, Data: Data}
}

func (frame CANFrame) WriteCANFrameToByteArray() []byte {
	bs := make([]byte, 28)
	id := make([]byte, 4)
	eId := make([]byte, 4)
	ide := make([]byte, 4)
	rtr := make([]byte, 4)
	dlc := make([]byte, 4)

	binary.BigEndian.PutUint32(id, frame.StdId)
	binary.BigEndian.PutUint32(eId, frame.ExtendedId)
	binary.BigEndian.PutUint32(ide, frame.IDE)
	binary.BigEndian.PutUint32(rtr, frame.RTR)
	binary.BigEndian.PutUint32(dlc, frame.DLC)

	copy(bs[0:4], id)
	copy(bs[4:8], eId)
	copy(bs[8:12], ide)
	copy(bs[12:16], rtr)
	copy(bs[16:20], dlc)
	copy(bs[20:28], frame.Data)
	return bs
}

func ReadCANFrameFromByteArray(bytes []byte) CANFrame {
	canFrame := new(CANFrame)

	id := binary.BigEndian.Uint32(bytes[0:4])
	eid := binary.BigEndian.Uint32(bytes[4:8])
	ide := binary.BigEndian.Uint32(bytes[8:12])
	rtr := binary.BigEndian.Uint32(bytes[12:16])
	dlc := binary.BigEndian.Uint32(bytes[16:20])
	data := bytes[20:28]

	canFrame.StdId = id
	canFrame.ExtendedId = eid
	canFrame.IDE = ide
	canFrame.RTR = rtr
	canFrame.DLC = dlc
	canFrame.Data = data

	return *canFrame
}
