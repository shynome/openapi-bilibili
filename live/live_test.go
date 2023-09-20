package live

import (
	"context"
	"log"
	"testing"

	"github.com/shynome/err0/try"
	bilibili "github.com/shynome/openapi-bilibili"
	"github.com/shynome/openapi-bilibili/internal/testutil"
)

func TestConnect(t *testing.T) {
	if !testutil.Debug {
		log.Println("skip")
		return
	}

	conf := testutil.Conf

	ctx := context.Background()

	c := bilibili.NewClient(conf.Key, conf.Secret)

	s := try.To1(c.Connect(ctx, conf.AppID, conf.IDCode))
	try.To(s.Close())
	// defer s.Close()
	// go s.Keepalive(ctx)

	info := s.Info().WebsocketInfo
	ch := try.To1(Connect(ctx, info.AuthBody, info.WssLink))
	for msg := range ch {
		log.Println("data", msg.Cmd, string(msg.Data))
	}
}
