package live

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	ph := NewPacketHeader(OpAuth, MsgV0, 114514)
	assert.Equal(t, ph.End(), uint32(114514))
	assert.Equal(t, ph.Operation(), OpAuth)
	assert.Equal(t, ph.Version(), MsgV0)
}
