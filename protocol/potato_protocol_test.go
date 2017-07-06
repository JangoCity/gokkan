package protocol

//
//import (
//	"github.com/michey/gokkan/data"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestProtocolDecodeCANFrame(t *testing.T) {
//	p := PotatoProtocol{}
//	bytes := []byte{'!', 0, 0, 0, 255, 0, 0, 0, 8, 0, 0, 0, 4, 0, 0, 0, 2, 0, 0, 0, 255, 1, 2, 3, 4, 5, 6, 7, 8, '\n'}
//	_, canFrame := p.Decode(bytes)
//
//	assert.Exactly(t, canFrame.StdId, uint32(255), "StdId should be zero!")
//	assert.Exactly(t, canFrame.ExtendedId, uint32(8), "ExtendedId should be zero!")
//	assert.Exactly(t, canFrame.IDE, uint32(4), "IDE should be zero!")
//	assert.Exactly(t, canFrame.RTR, uint32(2), "RTR should be zero!")
//	assert.Exactly(t, canFrame.DLC, uint32(255), "DLC should be zero!")
//	assert.Exactly(t, canFrame.Data, []byte{1, 2, 3, 4, 5, 6, 7, 8}, "Data should be zero!")
//}
//
//func TestProtocolEncodeCanFrame(t *testing.T) {
//	p := PotatoProtocol{}
//	canFrm := data.CANFrame{StdId: 0xff, ExtendedId: 0x00, IDE: data.CAN_IDE_STD, RTR: data.CAN_RTR_DATA, DLC: 8, Data: []byte{1, 2, 3, 4, 5, 6, 7, 8}}
//	bytes := p.Encode(&canFrm)
//	assert.Exactly(t, bytes, []byte{'>', 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 1, 2, 3, 4, 5, 6, 7, 8, '\n'}, "Encode blah-blah-blah")
//}
