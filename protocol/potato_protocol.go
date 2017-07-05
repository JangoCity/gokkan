package protocol

import (
	"encoding/binary"
	"fmt"
	"github.com/michey/gokkan/data"
)

const (
	OUTGOING_FRAME  uint8 = '>'
	INCOMING_FRAME  uint8 = '!'
	OUTGOING_FILTER uint8 = '^'
	STOP_CHAR             = '#'
)

type PotatoProtocol struct {
	IProtocol
}

func (p *PotatoProtocol) Encode(frame *data.CANFrame) []byte {
	bytes := make([]byte, 30)
	bytes[0] = OUTGOING_FRAME
	bytes[29] = STOP_CHAR

	data_bytes := bytes[1:30]
	copy(data_bytes, writeCANFrameToByteArray(frame))
	return bytes
}

func (p *PotatoProtocol) Decode(bytes []byte) (err error, frame *data.CANFrame) {
	fmt.Printf("%+v\n", bytes)
	return nil, readCANFrameFromByteArray(bytes[1:30])
}

func writeCANFrameToByteArray(frame *data.CANFrame) []byte {
	bs := make([]byte, 28)
	id := make([]byte, 4)
	eId := make([]byte, 4)
	ide := make([]byte, 4)
	rtr := make([]byte, 4)
	dlc := make([]byte, 4)

	binary.LittleEndian.PutUint32(id, frame.StdId)
	binary.LittleEndian.PutUint32(eId, frame.ExtendedId)
	binary.LittleEndian.PutUint32(ide, uint32(frame.IDE))
	binary.LittleEndian.PutUint32(rtr, uint32(frame.RTR))
	binary.LittleEndian.PutUint32(dlc, frame.DLC)

	copy(bs[0:4], id)
	copy(bs[4:8], eId)
	copy(bs[8:12], ide)
	copy(bs[12:16], rtr)
	copy(bs[16:20], dlc)
	copy(bs[20:28], frame.Data)
	return bs
}

func readCANFrameFromByteArray(bytes []byte) *data.CANFrame {
	canFrame := new(data.CANFrame)

	id := binary.LittleEndian.Uint32(bytes[0:4])
	eid := binary.LittleEndian.Uint32(bytes[4:8])
	ide := data.IDE(binary.LittleEndian.Uint32(bytes[8:12]))
	rtr := data.RTR(binary.LittleEndian.Uint32(bytes[12:16]))
	dlc := binary.LittleEndian.Uint32(bytes[16:20])
	data := bytes[20:28]

	canFrame.StdId = id
	canFrame.ExtendedId = eid
	canFrame.IDE = ide
	canFrame.RTR = rtr
	canFrame.DLC = dlc
	canFrame.Data = data

	return canFrame
}
