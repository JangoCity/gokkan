package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullFrame(t *testing.T) {
	cf := InitCanFrame(0x7ff, 8, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	f := cf.Decode();
	assert.Exactly(t, f, []byte("t7ff8ffffffffffffffff\r"), "Data")
}

func TestShrinkedFrame(t *testing.T) {
	cf2 := InitCanFrame(0x321, 5, []byte{0xff, 0xff, 0xff, 0xff, 0xff})
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("t3215ffffffffff\r"), "Data")
}

func TestZeroFrame(t *testing.T) {
	cf2 := InitCanFrame(0x001, 0, []byte{})
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("t0010\r"), "Data")
}

func TestExtendedShrinkedFrame(t *testing.T) {
	cf2 := InitExtCanFrame(0x1fffffff, 4, []byte{0x22, 0x23, 0x24, 0x25})
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("T1fffffff422232425\r"), "Data")
}

func TestExtendedFrame(t *testing.T) {
	cf2 := InitExtCanFrame(0x1fffffff, 4, []byte{0x22, 0x23, 0x24, 0x25})
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("T1fffffff422232425\r"), "Data")
}

func TestExtendedFrame2(t *testing.T) {
	cf2 := InitExtCanFrame(0x0000ffff, 7, []byte{0x3, 0x21, 0x4, 0x25, 0x6, 0x26, 0x7})
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("T0000ffff703210425062607\r"), "Data")
}

func TestRtrFrame(t *testing.T) {
	cf2 := InitRtrCanFrame(0x022, 7)
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("r0227\r"), "Data")
}

func TestExtRtrFrame(t *testing.T) {
	cf2 := InitRtrExtCanFrame(0x00ff00ff, 3)
	f2 := cf2.Decode();
	assert.Exactly(t, f2, []byte("R00ff00ff3\r"), "Data")
}
