package bilibili

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

func ApiCall[P any, T any](c *Client, api string) func(ctx context.Context, payload P) (_ T, err error) {
	return func(ctx context.Context, payload P) (_ T, err error) {
		defer err0.Then(&err, nil, func() {
			err = errors.Join(
				err,
				fmt.Errorf(
					"api: %s, payload: %+v",
					api, payload),
			)
		})
		body := try.To1(json.Marshal(payload))
		link := try.To1(url.JoinPath(c.Endpoint, api))
		req := try.To1(http.NewRequest(http.MethodPost, link, bytes.NewReader(body)))
		req = req.WithContext(ctx)
		req = try.To1(c.NewApiRequest(req))
		resp := try.To1(c.Client.Do(req))
		defer resp.Body.Close()
		rbody := try.To1(io.ReadAll(resp.Body))
		if code := resp.StatusCode; code != 200 {
			err = fmt.Errorf(
				"status code expect 200, but got %d. body: %s",
				code, string(rbody))
			return
		}
		var result Response[json.RawMessage]
		try.To(json.Unmarshal(rbody, &result))
		if result.Code != 0 {
			err = errors.Join(ErrBilibiliApiError, &result)
			return
		}
		var data T
		try.To(json.Unmarshal(result.Data, &data))
		return data, nil
	}
}

type Response[T any] struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Data      T      `json:"data"`
}

var _ error = (*Response[any])(nil)

func (r *Response[T]) Error() string {
	return fmt.Sprintf(
		"错误码 %d, 错误原因: %s. RequestId: %s",
		r.Code, r.Message, r.RequestId)
}

var ErrBilibiliApiError = errors.New("bilibili 接口返回报错")
