package cmd

type SuperChat struct {
	MsgBase
	UserInfo
	MessageID int64  `json:"message_id"` // 留言id(风控场景下撤回留言需要)
	Message   string `json:"message"`    // 留言内容
	Rmb       int64  `json:"rmb"`        // 支付金额(元)
	StartTime int64  `json:"start_time"` // 生效开始时间
	EndTime   int64  `json:"end_time"`   // 生效结束时间
}

type DelSuperChat struct {
	RoomID     int64   `json:"room_id"`     // 直播间id
	MessageIDs []int64 `json:"message_ids"` // 要撤回的留言id
	MsgID      string  `json:"msg_id"`      // 消息唯一id
}
