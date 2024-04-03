package cmd

type AnchorInfo struct {
	// 收礼主播uid
	//
	// Deprecated: B站会在 2024-04-15 把 UID 设为 0
	UID      int64  `json:"uid"`
	OpenID   string `json:"open_id"` // 收礼主播open_id
	Username string `json:"uname"`   // 收礼主播昵称
	Uface    string `json:"uface"`   // 收礼主播头像

}
