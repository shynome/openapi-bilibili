package bilibili

import (
	"net/url"
	"testing"

	"github.com/shynome/err0/try"
)

func TestVerifyH5Params(t *testing.T) {
	var link = "https://play-live.bilibili.com/?Timestamp=1650012983&Code=460803&Mid=110000345&Caller=bilibili&CodeSign=8c1fa83955d83960680277122bd31fd6f209a82787d57912c1d3817487bfc2ef"
	u := try.To1(url.Parse(link))
	bclient := NewClient("", "NPRZADNURSKNGYDFMDKJOOTLQMGDHL")
	try.To(bclient.VerifyH5Params(u.Query()))
	return
}
