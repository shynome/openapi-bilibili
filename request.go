package bilibili

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

func (c *Client) NewApiRequest(req *http.Request) (_ *http.Request, err error) {
	defer err0.Then(&err, nil, nil)
	if req.Method != http.MethodPost {
		return nil, ErrBilibiliApiSupportOnlyPost
	}

	body := try.To1(io.ReadAll(req.Body))
	req.Body = io.NopCloser(bytes.NewReader(body))

	bh := c.NewBaseHeader(body)
	req.Header = MergeHeader(bh, req.Header)
	signature := c.CreateSignature(req.Header)
	req.Header.Set("Authorization", signature)

	return req, nil
}

func (c *Client) NewBaseHeader(data []byte) http.Header {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	bodyMd5 := fmt.Sprintf("%x", md5.Sum(data))
	nonce := uuid.NewString()

	header := make(http.Header)
	// 加密要用到的header
	header.Set("x-bili-timestamp", timestamp)
	header.Set("x-bili-signature-method", "HMAC-SHA256")
	header.Set("x-bili-signature-nonce", nonce)
	header.Set("x-bili-accesskeyid", c.AccessKey)
	header.Set("x-bili-signature-version", "1.0")
	header.Set("x-bili-content-md5", bodyMd5)

	// 一些基础header
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")
	return header
}

var sortedSignatureHeaders = []string{
	"x-bili-timestamp",
	"x-bili-signature-method",
	"x-bili-signature-nonce",
	"x-bili-accesskeyid",
	"x-bili-signature-version",
	"x-bili-content-md5",
}

func init() {
	sort.Strings(sortedSignatureHeaders)
}

type SignatureHeader http.Header

const brLen = len("\n")

func (header SignatureHeader) ToBytes() []byte {
	buf := new(bytes.Buffer)
	h := http.Header(header)
	for _, k := range sortedSignatureHeaders {
		v := h.Get(k)
		s := k + ":" + v
		fmt.Fprintln(buf, s)
	}
	return buf.Bytes()[:buf.Len()-brLen]
}

func (c *Client) CreateSignature(header http.Header) string {
	data := SignatureHeader(header).ToBytes()
	return c.hmacSHA256(data)
}

func MergeHeader(dst, src http.Header) http.Header {
	dst = dst.Clone()
	for k := range src {
		dst.Set(k, src.Get(k))
	}
	return dst
}

func (c *Client) hmacSHA256(data []byte) string {
	mac := hmac.New(sha256.New, []byte(c.AccessKeySecret))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

var ErrBilibiliApiSupportOnlyPost = errors.New("bilibili 接口仅支持 POST 请求")
