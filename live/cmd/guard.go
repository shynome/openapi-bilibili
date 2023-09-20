package cmd

type Guard struct {
	MsgBase
	UserInfo AnchorInfo `json:"user_info"` // 用户信息
	GuardInfo
	FansMedal
}

type GuardInfo struct {
	Level GuardLevel `json:"guard_level"` // 大航海等级
	Num   int64      `json:"guard_num"`   // 大航海数量
	Unit  string     `json:"guard_unit"`  // 大航海单位
}

type GuardLevel int64

const (
	Lv1Guard GuardLevel = 1 + iota // 1: 总督
	Lv2Guard                       // 2: 提督
	Lv3Guard                       // 3: 舰长
)
