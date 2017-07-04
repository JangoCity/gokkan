package data

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

type Response struct {
	*CANFrame
	Timestamp int64
}
