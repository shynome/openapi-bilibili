package live_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/shynome/err0/try"
	bilibili "github.com/shynome/openapi-bilibili"
	"github.com/shynome/openapi-bilibili/internal/testutil"
	"github.com/shynome/openapi-bilibili/live"
)

func TestConnect(t *testing.T) {
	if !testutil.Debug {
		log.Println("skip")
		return
	}

	conf := testutil.Conf

	ctx := context.Background()

	c := bilibili.NewClient(conf.Key, conf.Secret)

	app := try.To1(c.Open(ctx, conf.AppID, conf.IDCode))
	time.AfterFunc(2*time.Minute, func() {
		app.Close()
	})
	go app.KeepAlive(ctx)

	room := live.RoomWith(app.Info().WebsocketInfo)
	ch := try.To1(room.Connect(ctx))
	for msg := range ch {
		log.Println("data", msg.Cmd, string(msg.Data))
	}
	t.Log("closed")
}
