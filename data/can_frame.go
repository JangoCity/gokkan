package data

import "fmt"

type IDE uint32

const (
	MIN_STD_ID = 0x00
	MAX_STD_ID = 0x7FF
)

//CAN Identifier Type
const (
	CAN_IDE_STD IDE = 0
	CAN_IDE_EXT IDE = 4
)

type RTR uint32

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

type Response struct {
	*CANFrame
	Timestamp int64
}

func (r *Response) String() string {
	return fmt.Sprintf("timestamp:%d; ", r.Timestamp) + r.CANFrame.String()
}
