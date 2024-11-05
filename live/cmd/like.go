package cmd

type Like struct {
	MsgBase // 文档上没写有 msg_id
	AnchorInfo
	LikeText  string `json:"like_text"`  // 点赞文案( “xxx点赞了”)
	LikeCount string `json:"like_count"` // 对单个用户最近2秒的点赞次数聚合
	FansMedal
}

type Enter struct {
	MsgBase
	AnchorInfo2

	// Deprecated: 此信息无此字段
	MsgId *string
}
