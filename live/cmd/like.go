package cmd

type Like struct {
	MsgBase // 文档上没写有 msg_id
	AnchorInfo
	LikeText string `json:"like_text"` // 点赞文案( “xxx点赞了”)
	FansMedal
}
