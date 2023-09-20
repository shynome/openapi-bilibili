package cmd

type CMD string

const (
	DanmuCMD        = "LIVE_OPEN_PLATFORM_DM"             // 获取弹幕信息
	GiftCMD         = "LIVE_OPEN_PLATFORM_SEND_GIFT"      // 获取礼物信息
	SuperChatCMD    = "LIVE_OPEN_PLATFORM_SUPER_CHAT"     // 获取付费留言
	DelSuperChatCMD = "LIVE_OPEN_PLATFORM_SUPER_CHAT_DEL" // 付费留言下线
	GuardCMD        = "LIVE_OPEN_PLATFORM_GUARD"          // 付费大航海
	LikeCMD         = "LIVE_OPEN_PLATFORM_LIKE"           // 点赞信息
)

type Cmd[T any] struct {
	CMD  CMD `json:"cmd"`
	Data T   `json:"data"`
}

type MsgBase struct {
	RoomID    int64  `json:"room_id"`   // 房间号
	Timestamp int64  `json:"timestamp"` // 时间戳
	MsgId     string `json:"msg_id"`    // 消息唯一id
}
