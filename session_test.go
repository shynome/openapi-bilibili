package bilibili

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/shynome/err0/try"
)

var testConf struct {
	Key    string
	Secret string
	AppID  int64
	IDCode string
}

func TestMain(m *testing.M) {
	testConf.Key = os.Getenv("BILIBILI_KEY")
	testConf.Secret = os.Getenv("BILIBILI_SECRET")
	appID := os.Getenv("BILIBILI_APPID")
	testConf.AppID = try.To1(strconv.ParseInt(appID, 10, 64))
	testConf.IDCode = os.Getenv("BILIBILI_IDCODE")
	m.Run()
}

func TestSession(t *testing.T) {
	ctx := context.Background()
	client := NewClient(testConf.Key, testConf.Secret)
	session := try.To1(client.Connect(ctx, testConf.AppID, testConf.IDCode))
	ctx2, cause := context.WithCancelCause(ctx)
	go func() {
		err := session.Keepalive(ctx)
		cause(err)
	}()
	defer func() {
		try.To(session.Close())
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
