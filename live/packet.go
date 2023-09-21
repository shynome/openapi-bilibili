package live

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"github.com/shynome/openapi-bilibili/live/cmd"
)

type Msg = cmd.Cmd[json.RawMessage]

func Unpack(ch chan<- Msg, data []byte) (err error) {
	defer err0.Then(&err, nil, nil)
	if l := len(data); l < 16 {
		return
	}
	h := PacketHeader(data[0:16])
	end := h.End()
	if realEnd := len(data); end > uint32(realEnd) {
		return fmt.Errorf(
			"%w. header length is %d, but the data length is %d", ErrMsgPackedWrong,
			end, realEnd)
	}
	body := data[16:end]
	switch v := h.Version(); v {
	default:
		return fmt.Errorf("unsupported msg version %v", v)
	case MsgV2:
		r := bytes.NewReader(body)
		zr := try.To1(zlib.NewReader(r))
		defer zr.Close()
		data = try.To1(io.ReadAll(zr))
		return Unpack(ch, data)
	case MsgV0:
		switch op := h.Operation(); op {
		case OpAuthReply:
			ch <- Msg{Cmd: "auth_reply", Data: body}
		case OpHeartbeatReply:
			ch <- Msg{Cmd: "heartbeat_reply", Data: body}
		default:
			ch <- Msg{Cmd: "unknown", Data: body}
		case OpSmsSendReply:
			var msg Msg
			try.To(json.Unmarshal(body, &msg))
			ch <- msg
		}
		return Unpack(ch, data[end:])
	}
}

var ErrMsgPackedWrong = errors.New("msg packed wrong")
