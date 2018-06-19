package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStateControlMSg(t *testing.T) {
	assert.Exactly(t, StateMsg(OpenState), []byte("O\r"))
	assert.Exactly(t, StateMsg(LoopbackState), []byte("L\r"))
	assert.Exactly(t, StateMsg(ResetState), []byte("R\r"))
}
