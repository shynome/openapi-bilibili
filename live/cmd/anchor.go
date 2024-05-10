package cmd

type AnchorInfo struct {
	// 收礼主播uid
	//
	// Deprecated: 已废弃，固定为0
	UID      int64  `json:"uid"`
	OpenID   string `json:"open_id"` // 收礼主播open_id
	Username string `json:"uname"`   // 收礼主播昵称
	Uface    string `json:"uface"`   // 收礼主播头像

}
