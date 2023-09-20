package testutil

import (
	"os"
	"strconv"

	"github.com/shynome/err0/try"
)

var Conf struct {
	Key    string
	Secret string
	AppID  int64
	IDCode string
}

func init() {
	Conf.Key = os.Getenv("BILIBILI_KEY")
	Conf.Secret = os.Getenv("BILIBILI_SECRET")
	appID := os.Getenv("BILIBILI_APPID")
	Conf.AppID = try.To1(strconv.ParseInt(appID, 10, 64))
	Conf.IDCode = os.Getenv("BILIBILI_IDCODE")
}
