package cmd

type Danmu struct {
	MsgBase
	UserInfo
	Msg   string    `json:"msg"`           // 弹幕内容
	Emoji string    `json:"emoji_img_url"` // 表情包图片地址
	Type  DanmuType `json:"dm_type"`       // 弹幕类型 0：普通弹幕 1：表情包弹幕
}

type DanmuType int64

const (
	DanmuNormal DanmuType = 0 // 0: 普通弹幕
	DanmuEmoji  DanmuType = 1 // 1：表情包弹幕
)
