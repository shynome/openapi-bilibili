# Changelog

## [0.4.3] - 2024-05-10

- 添加 B 站新增的两个字段 `Like.LikeCount` 和 `Guard.Price`

## [0.4.2] - 2024-04-26

- 支持最新的消息推送结束通知

## [0.4.1] - 2024-04-14

- 添加 `room.WSDialOptioins` 以便支持代理

## [0.4.0] - 2024-04-11

- 添加对野开弹幕字段 `info` 的支持

## [0.3.1] - 2024-04-03

- 把 UID 字段加回来了

## [0.3.0] - 2024-03-12

- follow bilibili change, replace `uid` with `open_id`

## [0.2.0] - 2024-02-15

- 添加批量心跳接口

## [0.1.0] - 2023-11-10

- add: 添加 h5 参数校验

## [0.0.11] - 2023-11-07

- fix: handle try.To err

## [0.0.10] - 2023-09-23

### Improve

- 为 `GuardLevel` 实现 `fmt.Stringer`

### Fix

- 修复礼物数量类型错误

## [0.0.8] - 2023-09-21

### Addd

- 添加 `bilibili.Time`, 方便解析时间

## [0.0.7] - 2023-09-21

众多修复, 建议立即升级

### Fix

- 修复消息解码
- 修复可能的[]byte 越界访问错误
- 修复 auth write 未设置超时
- 修复 ping timer 未按照代码预期停止

## [0.0.6] - 2023-09-21

### Fix

-修复了一些类型错误

## [0.0.5] - 2023-09-21

### Fix

- 添加 Ping-Pong 检查 WebSocket 存活状态

## [0.0.4] - 2023-09-21

### Improve

- 导出 live.ErrAuthFailed 错误
- 添加超时机制

## [0.0.3] - 2023-09-21

### Add

- 可设定 bilibili open-live 接口地址
