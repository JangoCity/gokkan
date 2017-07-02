package protocol

import "github.com/michey/gokkan/data"

const (
	OUTGOING_FRAME uint8 = '>'
	INCOMING_FRAME uint8 = '!'
	OUTGOING_FILTE uint8 = '^'
)

func Marshall(frame data.CANFrame) []byte {
	bytes := make([]byte, 30)
	bytes[0] = OUTGOING_FRAME
	bytes[29] = '\n'

	data_bytes := bytes[1:30]
	copy(data_bytes, frame.WriteCANFrameToByteArray())
	return bytes
}

func readString(bytes []byte) {
	//if bytes[1] == INCOMING_FRAME {
	//	canFrame := parseCanFrame(bytes)
	//}
}

func Unarshall(bytes []byte) data.CANFrame {
	return data.ReadCANFrameFromByteArray(bytes[1:30])
}
