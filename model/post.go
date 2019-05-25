package model

const (
	POST_SYSTEM_MESSAGE_PREFIX = "system_"
	POST_DEFAULT               = ""
	POST_SLACK_ATTCHMENT       = "slack_attachment"
	POST_SYSTEM_GENERIC        = "system_generic"
	POST_JOIN_LEAVE            = "system_join_leave"
	POST_JOIN_CHANNEL          = "system_join_channel"
	POST_LEAVE_CHANNEL         = "system_leave_channel"
)

type Post struct {
	Id         string `json:"id"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	EditAt     int64  `json:"edit_at"`
	DeleteAt   int64  `json:"delete_at"`
	IsPinned   bool   `json:"is_pinned"`
	UserId     string `json:"user_id"`
	ChannelId  string `json:"channel_id"`
	RootId     string `json:"root_id"`
	ParentId   string `json:"parent_id"`
	OriginalId string `json:"original_id"`

	Message string `json:"message"`

	// MessageSource will contain the message as submitted by the user if Message has been modified
	// by Mattermost for presentation (e.g. if an image proxy is being used). It should be used to
	// populate edit boxes if present.
	MessageSource string `json:"message_source,omitempty" db:"-"`

	Type          string          `json:"type"`
	Props         StringInterface `json:"props"`
	Hashtags      string          `json:"hashtags"`
	Filenames     StringArray     `json:"filenames,omitempty"`
	FileIds       StringArray     `json:"file_ids,omitempty"`
	PendingPostId string          `json:"pending_post_id" db:"-"`
	HasReactions  bool            `json:"has_reactions,omitempty"`
}
