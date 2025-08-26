package live

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	bilibili "github.com/shynome/openapi-bilibili"
	"nhooyr.io/websocket"
)

type Room struct {
	info    bilibili.WebsocketInfo
	game_id string

	WSDialOptioins *websocket.DialOptions
}

func RoomWith(info bilibili.WebsocketInfo, game_id string) *Room {
	return &Room{
		info:    info,
		game_id: game_id,
	}
}

func (room *Room) Connect(ctx context.Context) (_ <-chan Msg, err error) {
	defer err0.Then(&err, nil, nil)

	ctx, cause := context.WithCancelCause(ctx)
	defer err0.Then(&err, nil, func() {
		cause(err)
	})

	info := room.info
	var conn *websocket.Conn
	{
		errs := []error{}
		for _, link := range info.WssLink {
			c, _, err := websocket.Dial(ctx, link, room.WSDialOptioins)
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

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

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
			return nil, fmt.Errorf("%w. code %d", ErrAuthFailed, status.Code)
		}
	}

	ch := make(chan Msg, 1024)

	go func() {
		defer close(ch)
		var wg sync.WaitGroup
		defer wg.Wait()
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
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := Unpack(ch, msg); err != nil {
					if end, ok := err.(*GameEnd); ok && end.GameID == room.game_id {
						cause(err)
						return
					}
					ch <- Msg{
						Cmd:  "unpack error",
						Data: json.RawMessage(err.Error()),
					}
				}
			}()
		}
	}()

	go func() {
		timer := time.NewTicker(5 * time.Second)
		defer timer.Stop()
		ping := func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return conn.Ping(ctx)
		}
		for range timer.C {
			if err := ping(ctx); err != nil {
				cause(err)
				return
			}
		}
	}()

	go func() {
		defer cause(nil)

		hdr := NewPacketHeader(OpHeartbeat, MsgV0, 16)

		timer := time.NewTicker(20 * time.Second)
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

var ErrAuthFailed = errors.New("auth failed")
