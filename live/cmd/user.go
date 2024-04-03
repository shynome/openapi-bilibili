package cmd

type UserInfo struct {
	// 用户uid
	//
	// Deprecated: B站会在 2024-04-15 把 UID 设为 0
	UID        int64  `json:"uid"`
	Username   string `json:"uname"`       // 用户昵称
	OpenID     string `json:"open_id"`     // 用户open_id
	Uface      string `json:"uface"`       // 用户头像
	GuardLevel int64  `json:"guard_level"` // 对应房间大航海等级
	FansMedal
}

type FansMedal struct {
	WearingStatus bool   `json:"fans_medal_wearing_status"` // 该房间粉丝勋章佩戴情况
	Name          string `json:"fans_medal_name"`           // 粉丝勋章名
	Level         int64  `json:"fans_medal_level"`          // 对应房间勋章信息
}
