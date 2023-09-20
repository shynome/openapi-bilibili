package cmd

type AnchorInfo struct {
	UID      int64  `json:"uid"`   // 收礼主播uid
	Username string `json:"uname"` // 收礼主播昵称
	Uface    string `json:"uface"` // 收礼主播头像
}
