package model

const (
	ME                  = "me"
	USER_NOTIFY_ALL     = "all"
	USER_NOTIFY_MENTION = "mention"

	DEFAULT_LOCALE          = "en"
	USER_AUTH_SERVICE_EMAIL = "email"

	USER_EMAIL_MAX_LENGTH     = 128
	USER_NICKNAME_MAX_RUNES   = 64
	USER_POSITION_MAX_RUNES   = 128
	USER_FIRST_NAME_MAX_RUNES = 64
	USER_LAST_NAME_MAX_RUNES  = 64
	USER_AUTH_DATA_MAX_LENGTH = 128
	USER_NAME_MAX_LENGTH      = 64
	USER_NAME_MIN_LENGTH      = 64
	USER_PASSWORD_MAX_LENGTH  = 72
	USER_LOCALE_MAX_LENGTH    = 5
)

type User struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type UserPatch struct {
	Username *string `json:"username"`
}

type UserAuth struct {
	Password string  `json:"password,omitempty"`
	AuthData *string `json:"auth_data,omitempty"`
}
