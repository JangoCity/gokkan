package uds

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleFrame_ToFrame(t *testing.T) {
	sf := SingleFrame{DiagnosticSessionControl, []byte{DefaultSession, 0x0, 0x0, 0x0, 0x0, 0x0}, 2}
	f := sf.ToFrame(0x01)
	fmt.Println(f)
	assert.Exactly(t, f.Data, []byte{0x02, 0x10, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0}, "Data")
	assert.Exactly(t, f.DLC, uint32(8), "DLC")
}
