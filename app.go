package bilibili

import (
	"context"
	"encoding/json"
	"time"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

type App struct {
	AppOpen
	Client *Client
	info   *AppInfo

	timerStop context.CancelFunc
}

type AppOpen struct {
	AppID  int64  `json:"app_id"`
	IDCode string `json:"code"` //主播身份码
}

type AppInfo struct {
	// 场次信息
	GameInfo struct {
		GameId string `json:"game_id"`
	} `json:"game_info"`
	// 长连信息
	WebsocketInfo WebsocketInfo `json:"websocket_info"`
}

type WebsocketInfo struct {
	//  长连使用的请求json体 第三方无需关注内容,建立长连时使用即可
	AuthBody string `json:"auth_body"`
	//  wss 长连地址
	WssLink []string `json:"wss_link"`
}

func (c *Client) Open(ctx context.Context, appid int64, code string) (_ *App, err error) {
	defer err0.Then(&err, nil, nil)
	s := &App{
		Client: c,
		AppOpen: AppOpen{
			AppID:  appid,
			IDCode: code,
		},
	}
	start := ApiCall[AppOpen, *AppInfo](c, "/v2/app/start")
	s.info = try.To1(start(ctx, s.AppOpen))
	return s, nil
}

func (s *App) Info() *AppInfo {
	return s.info
}

type AppClose struct {
	// 场次id
	GameId string `json:"game_id"`
	// 项目id
	AppId int64 `json:"app_id"`
}

func (s *App) Close() (err error) {
	if s.timerStop != nil {
		s.timerStop()
	}
	end := ApiCall[AppClose, json.RawMessage](s.Client, "/v2/app/end")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err = end(ctx, AppClose{
		s.info.GameInfo.GameId,
		s.AppID,
	})
	return err
}

type AppKeepAlive struct {
	GameId string `json:"game_id"`
}

func (s *App) KeepAlive(ctx context.Context) error {
	if s.timerStop != nil {
		s.timerStop()
	}
	ctx, s.timerStop = context.WithCancel(ctx)
	keep := ApiCall[AppKeepAlive, json.RawMessage](s.Client, "/v2/app/heartbeat")
	timer := time.NewTicker(20 * time.Second)
	defer timer.Stop()
	gameId := s.info.GameInfo.GameId
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			_, err := keep(ctx, AppKeepAlive{GameId: gameId})
			if err != nil {
				return err
			}
		}
	}
}
