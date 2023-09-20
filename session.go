package bilibili

import (
	"context"
	"encoding/json"
	"time"

	"github.com/shynome/err0/try"
)

type Session struct {
	SessionConnect
	Client *Client
	info   *SessionInfo

	timerStop context.CancelFunc
}

type SessionConnect struct {
	AppID  int64  `json:"app_id"`
	IDCode string `json:"code"` //主播身份码
}

type SessionInfo struct {
	// 场次信息
	GameInfo struct {
		GameId string `json:"game_id"`
	} `json:"game_info"`
	// 长连信息
	WebsocketInfo struct {
		//  长连使用的请求json体 第三方无需关注内容,建立长连时使用即可
		AuthBody string `json:"auth_body"`
		//  wss 长连地址
		WssLink []string `json:"wss_link"`
	} `json:"websocket_info"`
}

func (c *Client) Connect(ctx context.Context, appid int64, code string) (*Session, error) {
	s := &Session{
		Client: c,
		SessionConnect: SessionConnect{
			AppID:  appid,
			IDCode: code,
		},
	}
	start := ApiCall[SessionConnect, *SessionInfo](c, "/v2/app/start")
	s.info = try.To1(start(ctx, s.SessionConnect))
	return s, nil
}

func (s *Session) Info() *SessionInfo {
	return s.info
}

type SessionClose struct {
	// 场次id
	GameId string `json:"game_id"`
	// 项目id
	AppId int64 `json:"app_id"`
}

func (s *Session) Close() (err error) {
	if s.timerStop != nil {
		s.timerStop()
	}
	end := ApiCall[SessionClose, json.RawMessage](s.Client, "/v2/app/end")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err = end(ctx, SessionClose{
		s.info.GameInfo.GameId,
		s.AppID,
	})
	return err
}

type SessionKeepAlive struct {
	GameId string `json:"game_id"`
}

func (s *Session) Keepalive(ctx context.Context) error {
	if s.timerStop != nil {
		s.timerStop()
	}
	ctx, s.timerStop = context.WithCancel(ctx)
	keep := ApiCall[SessionKeepAlive, json.RawMessage](s.Client, "/v2/app/heartbeat")
	timer := time.NewTicker(20 * time.Second)
	defer timer.Stop()
	gameId := s.info.GameInfo.GameId
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			_, err := keep(ctx, SessionKeepAlive{GameId: gameId})
			if err != nil {
				return err
			}
		}
	}
}
