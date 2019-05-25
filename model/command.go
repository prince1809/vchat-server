package model

const (
	COMMAND_METHOD_POST = "P"
	COMMAND_METHOD_GET  = "G"
	MIN_TRIGGER_LENGTH  = 1
	MAX_TRIGGER_LENGTH  = 128
)

type Command struct {
	Id               string `json:"id"`
	Token            string `json:"token"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	DeletedAt        int64  `json:"deleted_at"`
	CreatorId        string `json:"creator_id"`
	TeamId           string `json:"team_id"`
	Trigger          string `json:"trigger"`
	Method           string `json:"method"`
	Username         string `json:"username"`
	IconURL          string `json:"icon_url"`
	AutoComplete     bool   `json:"auto_complete"`
	AutoCompleteDesc string `json:"auto_complete_desc"`
	AutoCompleteHint string `json:"auto_complete_hint"`
	DisplayName      string `json:"display_name"`
	Description      string `json:"description"`
	URL              string `json:"url"`
}
