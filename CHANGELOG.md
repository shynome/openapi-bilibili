# Changelog

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
