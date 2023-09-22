package cmd

type Gift struct {
	MsgBase
	UserInfo
	GiftInfo
	Price      int64      `json:"price"`       // 支付金额(1000 = 1元 = 10电池),盲盒:爆出道具的价值
	Paid       bool       `json:"paid"`        // 是否是付费道具
	AnchorInfo AnchorInfo `json:"anchor_info"` // 主播信息
	ComboGift  bool       `json:"combo_gift"`  // 是否是combo道具
	ComboInfo  ComboInfo  `json:"combo_info"`  // 连击信息
}

type GiftInfo struct {
	ID   int64  `json:"gift_id"`   // 道具id(盲盒:爆出道具id)
	Name string `json:"gift_name"` // 道具名(盲盒:爆出道具名)
	Num  int64  `json:"gift_num"`  // 赠送道具数量
	Icon string `json:"gift_icon"` // 道具icon
}

type ComboInfo struct {
	BaseNum int64  `json:"combo_base_num"` // 每次连击赠送的道具数量
	Count   int64  `json:"combo_count"`    // 连击次数
	ID      string `json:"combo_id"`       // 连击id
	Timeout int64  `json:"combo_timeout"`  //连击有效期秒
}
