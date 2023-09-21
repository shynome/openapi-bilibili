package cmd

type CmdType string

const (
	CmdDanmu        CmdType = "LIVE_OPEN_PLATFORM_DM"             // 获取弹幕信息
	CmdGift         CmdType = "LIVE_OPEN_PLATFORM_SEND_GIFT"      // 获取礼物信息
	CmdSuperChat    CmdType = "LIVE_OPEN_PLATFORM_SUPER_CHAT"     // 获取付费留言
	CmdDelSuperChat CmdType = "LIVE_OPEN_PLATFORM_SUPER_CHAT_DEL" // 付费留言下线
	CmdGuard        CmdType = "LIVE_OPEN_PLATFORM_GUARD"          // 付费大航海
	CmdLike         CmdType = "LIVE_OPEN_PLATFORM_LIKE"           // 点赞信息
)

type Cmd[T any] struct {
	Cmd  CmdType `json:"cmd"`
	Data T       `json:"data"`
}

type MsgBase struct {
	RoomID    int64  `json:"room_id"`   // 房间号
	Timestamp int64  `json:"timestamp"` // 时间戳
	MsgId     string `json:"msg_id"`    // 消息唯一id
}
