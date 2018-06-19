package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpeedMessages(t *testing.T) {
	assert.Exactly(t, SetSpeedMsg(S10K), []byte("S0\r"))
	assert.Exactly(t, SetSpeedMsg(S20K), []byte("S1\r"))
	assert.Exactly(t, SetSpeedMsg(S50K), []byte("S2\r"))
	assert.Exactly(t, SetSpeedMsg(S100K), []byte("S3\r"))
	assert.Exactly(t, SetSpeedMsg(S125K), []byte("S4\r"))
	assert.Exactly(t, SetSpeedMsg(S250K), []byte("S5\r"))
	assert.Exactly(t, SetSpeedMsg(S500K), []byte("S6\r"))
	assert.Exactly(t, SetSpeedMsg(S800K), []byte("S7\r"))
	assert.Exactly(t, SetSpeedMsg(S1M), []byte("S8\r"))
	assert.Exactly(t, SetSpeedMsg(S83K), []byte("S9\r"))
}
