package model

type WebsocketBroadcast struct {
	OmitUsers             map[string]bool `json:"omit_users"`
	UserId                string          `json:"user_id"`
	ChannelId             string          `json:"channel_id"`
	TeamId                string          `json:"team_id"`
	ContainsSanitizedData bool            `json:"-"`
	ContainsSensitiveData bool            `json:"-"`
}

type WebSocketEvent struct {
	Event     string                 `json:"event"`
	Data      map[string]interface{} `json:"data"`
	Broadcast *WebsocketBroadcast    `json:"broadcast"`
	Sequence  int64                  `json:"seq"`
}
