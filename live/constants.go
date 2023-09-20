package live

type MsgVersion uint16 // Msg Version

const (
	MsgV0 MsgVersion = iota // 0 如果Version=0，Body中就是实际发送的数据。
	_                       //
	MsgV2                   // 2 如果Version=2，Body中是经过压缩后的数据，请使用zlib解压，然后按照Proto协议去解析
)

type Op uint32

const (
	OpHeartbeat      Op = 2 + iota // 2 客户端发送的心跳包(30秒发送一次)
	OpHeartbeatReply               // 3 服务器收到心跳包的回复
	_                              //
	OpSmsSendReply                 // 5 服务器推送的弹幕消息包
	_                              //
	OpAuth                         // 7 客户端发送的鉴权包(客户端发送的第一个包)
	OpAuthReply                    // 8 服务器收到鉴权包后的回复
)
