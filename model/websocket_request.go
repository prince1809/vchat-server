package model

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
)

type WebSocketRequest struct {
	// Client-provided fields
	Seq    int64                  `json:"seq"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`

	// Server-provided fields
	Session Session              `json:"-"`
	T       goi18n.TranslateFunc `json:"-"`
	Locale  string               `json:"-"`
}
