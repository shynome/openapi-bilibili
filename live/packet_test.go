package live

import (
	"bytes"
	"compress/zlib"
	"io"
	"testing"

	"github.com/shynome/err0/try"
	"github.com/stretchr/testify/assert"
)

func TestMsgV0(t *testing.T) {
	body := "hello world"

	var pkt = new(bytes.Buffer)
	hdr := NewPacketHeader(OpHeartbeatReply, MsgV0, uint32(16+len(body)))
	pkt.Write(hdr[:])
	pkt.WriteString(body)

	ch := make(chan Msg, 1)
	try.To(Unpack(ch, pkt.Bytes()))
	msg := <-ch
	assert.Equal(t, body, string(msg.Data))
}

func TestMsgV2(t *testing.T) {
	body := "hello world"

	body2 := func() []byte {
		buf := new(bytes.Buffer)
		w := zlib.NewWriter(buf)
		hdr := NewPacketHeader(OpHeartbeatReply, MsgV0, uint32(16+len(body)))
		w.Write(hdr[:])
		io.WriteString(w, body)
		w.Close()
		return buf.Bytes()
	}()

	pkt2 := new(bytes.Buffer)
	hdr2 := NewPacketHeader(OpHeartbeatReply, MsgV2, uint32(16+len(body2)))
	pkt2.Write(hdr2[:])
	pkt2.Write(body2)

	ch := make(chan Msg, 1)
	try.To(Unpack(ch, pkt2.Bytes()))
	msg := <-ch
	assert.Equal(t, body, string(msg.Data))
}
