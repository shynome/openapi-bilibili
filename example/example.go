package main

import (
	"context"
	"log"
	"time"

	"github.com/shynome/err0/try"
	bilibili "github.com/shynome/openapi-bilibili"
	"github.com/shynome/openapi-bilibili/live"
)

func main() {
	bclient := bilibili.NewClient("access_key_id", "access_key_secret")
	ctx := context.Background()
	app := try.To1(bclient.Open(ctx, 0000000000000, "主播身份码"))
	try.To(app.Close()) // 拿到认证信息就可以关闭了
	info := app.Info().WebsocketInfo
	room := live.RoomWith(info)
	ctx, closeMsgCh := context.WithCancel(ctx)
	go func() {
		time.Sleep(10 * time.Minute)
		closeMsgCh()
	}()
	msgCh := try.To1(room.Connect(ctx))
	for msg := range msgCh {
		log.Println("msg", msg.Cmd, msg.Data)
	}
}
