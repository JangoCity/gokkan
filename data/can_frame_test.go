package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeZeroedCANFrame(t *testing.T) {
	bytes := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	canFrame := ReadCANFrameFromByteArray(bytes)
	assert.Exactly(t, canFrame.DLC, uint32(0), "DLC should be zero!")
	assert.Exactly(t, canFrame.StdId, uint32(0), "StdId should be zero!")
	assert.Exactly(t, canFrame.ExtendedId, uint32(0), "ExtendedId should be zero!")
	assert.Exactly(t, canFrame.IDE, uint32(0), "IDE should be zero!")
	assert.Exactly(t, canFrame.RTR, uint32(0), "RTR should be zero!")
	assert.Exactly(t, canFrame.Data, []byte{0, 0, 0, 0, 0, 0, 0, 0}, "Data should be zero!")
}

func TestDecodeDataCANFrame(t *testing.T) {
	bytes := []byte{0, 0, 0, 255, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0, 255, 1, 2, 3, 4, 5, 6, 7, 8}
	canFrame := ReadCANFrameFromByteArray(bytes)

	assert.Exactly(t, canFrame.StdId, uint32(255), "StdId should be zero!")
	assert.Exactly(t, canFrame.ExtendedId, uint32(8), "ExtendedId should be zero!")
	assert.Exactly(t, canFrame.IDE, uint32(4), "IDE should be zero!")
	assert.Exactly(t, canFrame.RTR, uint32(2), "RTR should be zero!")
	assert.Exactly(t, canFrame.DLC, uint32(255), "DLC should be zero!")
	assert.Exactly(t, canFrame.Data, []byte{1, 2, 3, 4, 5, 6, 7, 8}, "Data should be zero!")
}

func TestEncodeCanFrame(t *testing.T) {
	canFrm := CANFrame{StdId: 0xff, ExtendedId: 0x00, IDE: CAN_IDE_STD, RTR: CAN_RTR_DATA, DLC: 8, Data: []byte{1, 2, 3, 4, 5, 6, 7, 8}}
	bytes := canFrm.WriteCANFrameToByteArray()
	assert.Exactly(t, bytes, []byte{0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 1, 2, 3, 4, 5, 6, 7, 8}, "Encode blah-blah-blah")
}
