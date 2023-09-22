package bilibili

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/shynome/err0/try"
	"github.com/stretchr/testify/assert"
)

func TestTimeParse(t *testing.T) {
	var m struct {
		Time Time `json:"time"`
	}
	var body = `{"time": "2023-09-22 10:13:29"}`
	try.To(json.Unmarshal([]byte(body), &m))
	timestamp := time.Time(m.Time).Unix()
	assert.Equal(t, int64(1695348809), timestamp)
}
