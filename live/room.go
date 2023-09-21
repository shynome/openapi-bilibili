package live

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	bilibili "github.com/shynome/openapi-bilibili"
	"nhooyr.io/websocket"
)

type Room struct {
	info bilibili.WebsocketInfo
}

func RoomWith(info bilibili.WebsocketInfo) *Room {
	return &Room{
		info: info,
	}
}

func (room *Room) Connect(ctx context.Context) (_ <-chan Msg, err error) {
	defer err0.Then(&err, nil, nil)

	ctx, cause := context.WithCancelCause(ctx)

	info := room.info
	var conn *websocket.Conn
	{
		errs := []error{}
		for _, link := range info.WssLink {
			c, _, err := websocket.Dial(ctx, link, nil)
			if err == nil {
				conn = c
				break
			}
			errs = append(errs, err)
		}
		if conn == nil {
			return nil, errors.Join(errs...)
		}
	}

	{
		auth := new(bytes.Buffer)
		hdr := NewPacketHeader(OpAuth, MsgV0, uint32(len(info.AuthBody)+16))
		auth.Write(hdr[:])
		auth.WriteString(info.AuthBody)

		try.To(conn.Write(ctx, websocket.MessageBinary, auth.Bytes()))
		_, data := try.To2(conn.Read(ctx))
		ch := make(chan Msg, 1)
		try.To(Unpack(ch, data))
		msg := <-ch
		var status struct {
			Code int64 `json:"code"`
		}
		try.To(json.Unmarshal(msg.Data, &status))
		if status.Code != 0 {
			return nil, fmt.Errorf("auth fail. code %d", status.Code)
		}
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