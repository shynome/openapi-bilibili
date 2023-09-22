package bilibili

import (
	"strings"
	"time"
)

type Time time.Time

const TimeLayout = "2006-01-02 15:04:05"
const TimeZero = "0000-00-00 00:00:00"
const TimeOffset = 8 * 60 * 60

var TimeLocation *time.Location = time.FixedZone("Asia/Shanghai", TimeOffset)

func (c *Time) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" || value == TimeZero {
		return nil
	}

	t, err := time.ParseInLocation(TimeLayout, value, TimeLocation) //parse time
	if err != nil {
		return err
	}
	*c = Time(t) //set result using the pointer
	return nil
}
