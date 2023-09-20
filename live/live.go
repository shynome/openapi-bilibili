package live

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"nhooyr.io/websocket"
)

type Live struct{}

func Connect(ctx context.Context, authBody string, servers []string) (_ <-chan Msg, err error) {
	defer err0.Then(&err, nil, nil)

	ctx, cause := context.WithCancelCause(ctx)

	conn, _ := try.To2(websocket.Dial(ctx, servers[0], nil))

	{
		auth := new(bytes.Buffer)
		hdr := NewPacketHeader(OpAuth, MsgV0, uint32(len(authBody)+16))
		auth.Write(hdr[:])
		auth.WriteString(authBody)

		try.To(conn.Write(ctx, websocket.MessageBinary, auth.Bytes()))
	}

	ch := make(chan Msg, 1024)

	go func() {
		defer close(ch)
		for {
			_, msg, err := conn.Read(ctx)
			if err != nil {
				ch <- Msg{
					Cmd:  "error",
					Data: json.RawMessage(err.Error()),
				}
				cause(err)
				return
			}
			go func() {
				if err := Unpack(ch, msg); err != nil {
					log.Println("unpack error", err)
				}
			}()
		}
	}()

	go func() {
		defer cause(nil)

		hdr := NewPacketHeader(OpHeartbeat, MsgV0, 16)

		timer := time.NewTicker(30 * time.Second)
		defer timer.Stop()

		for range timer.C {
			heartbeat := new(bytes.Buffer)
			heartbeat.Write(hdr[:])
			if err := conn.Write(ctx, websocket.MessageBinary, heartbeat.Bytes()); err != nil {
				cause(err)
				return
			}
		}
	}()

	return ch, nil
}
