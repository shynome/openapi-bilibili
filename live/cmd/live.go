package cmd

type LiveStart struct {
	RoomID    int64  `json:"room_id"`   // 房间号
	OpenID    string `json:"open_id"`   // 主播open_id
	Timestamp int64  `json:"timestamp"` // 时间戳
}

type LiveEnd LiveStart
