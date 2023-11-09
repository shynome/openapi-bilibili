package bilibili

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"strings"
)

var sortedH5ParamKeys = []string{"Caller", "Code", "Mid", "Timestamp"}

func (client *Client) VerifyH5Params(q url.Values) error {
	d := ""
	codeSign := q.Get("CodeSign")
	if codeSign == "" {
		return fmt.Errorf("lost CodeSign")
	}
	for _, k := range sortedH5ParamKeys {
		v := q.Get(k)
		if v == "" {
			return fmt.Errorf("lost param: %s", k)
		}
		d += fmt.Sprintf("%s:%s\n", k, v)
	}
	d = strings.TrimSuffix(d, "\n")
	h := hmac.New(sha256.New, []byte(client.AccessKeySecret))
	io.WriteString(h, d)
	digest := h.Sum(nil)
	sign := hex.EncodeToString(digest)
	if sign != codeSign {
		return fmt.Errorf("CodeSign has wrong")
	}
	return nil
}
