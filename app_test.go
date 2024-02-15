package bilibili

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shynome/err0/try"
	"github.com/shynome/openapi-bilibili/internal/testutil"
)

func TestApp(t *testing.T) {
	conf := testutil.Conf

	ctx := context.Background()
	client := NewClient(conf.Key, conf.Secret)
	app := try.To1(client.Open(ctx, conf.AppID, conf.IDCode))
	ctx2, cause := context.WithCancelCause(ctx)
	go func() {
		err := app.KeepAlive(ctx)
		cause(err)
	}()
	defer func() {
		try.To(app.Close())
	}()
	defer func() {
		<-ctx2.Done()
		err := context.Cause(ctx2)
		if errors.Is(err, context.Canceled) {
			err = nil
		}
		if err != nil {
			t.Error(err)
		}
	}()
	{ // 测试 batch keep alive
		info := try.To1(client.BatchKeepAlive(ctx, []string{app.Info().GameInfo.GameId}))
		if len(info.FailedGameIDs) != 0 {
			t.Error("批量心跳失败")
		}
	}
	for {
		select {
		case <-time.After(2 * time.Minute):
			cause(nil)
			return
		case <-ctx2.Done():
			return
		}
	}
}
