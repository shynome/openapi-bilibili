package bilibili

import (
	"context"
	"os"
	"strconv"
	"testing"

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
	defer func() {
		try.To(session.Close())
	}()
}
