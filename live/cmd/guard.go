package cmd

import (
	"fmt"
	"strconv"
)

type Guard struct {
	MsgBase
	UserInfo AnchorInfo `json:"user_info"` // 用户信息
	GuardInfo
	FansMedal
}

type GuardInfo struct {
	Level GuardLevel `json:"guard_level"` // 大航海等级
	Num   int64      `json:"guard_num"`   // 大航海数量
	Unit  string     `json:"guard_unit"`  // 大航海单位(正常单位为“月”，如为其他内容，无视guard_num以本字段内容为准，例如*3天)
	Price int64      `json:"price"`       // 大航海金瓜子
}

type GuardLevel int64

const (
	GuardLv1 GuardLevel = 1 + iota // 1: 总督
	GuardLv2                       // 2: 提督
	GuardLv3                       // 3: 舰长
)

var _ fmt.Stringer = (*GuardLevel)(nil)

func (lv GuardLevel) String() string {
	switch lv {
	default:
		lv := strconv.FormatInt(int64(lv), 10)
		return fmt.Sprintf("未知级别: %s", lv)
	case GuardLv1:
		return "总督"
	case GuardLv2:
		return "提督"
	case GuardLv3:
		return "舰长"
	}
}
