package model

import "regexp"

var EMOJI_PATTERN = regexp.MustCompile(`:[a-zA-Z0-9_-]+:`)

type Emoji struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	CreatorId string `json:"creator_id"`
	Name      string `json:"name"`
}
