package model

import (
	"bytes"
	"encoding/base32"
	goi18n "github.com/mattermost/go-i18n/i18n"
	"github.com/pborman/uuid"
	"time"
)

const (
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
)

type StringInterface map[string]interface{}
type StringMap map[string]string
type StringArray []string

func (sa StringArray) Equals(input StringArray) bool {
	if len(sa) != len(input) {
		return false
	}

	for index := range sa {

		if sa[index] != input[index] {
			return false
		}
	}

	return true
}

var translateFunc goi18n.TranslateFunc = nil

func AppErrorInit(t goi18n.TranslateFunc) {
	translateFunc = t
}

type AppError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`        // Message to display to the end user without debugging information
	DetailedError string `json:"detailed_error"` // Internal error string to help the developer
	RequestId     string `json:"request_id"`     // The requestId that's also in the header
	StatusCode    int    `json:"status_code"`    // The http status code
	Where         string `json:"where"`          // The function where it happened in the form of struct.Func
	IsOAuth       bool   `json:"is_oauth"`       // Whether the error is OAuth specific
	params        map[string]interface{}
}

func (er *AppError) Error() string {
	return er.Where + ":" + er.Message + ", " + er.DetailedError
}

func (er *AppError) Translate(T goi18n.TranslateFunc) {
	if T == nil {
		er.Message = er.Id
		return
	}
	if er.params == nil {
		er.Message = T(er.Id)
	} else {
		er.Message = T(er.Id, er.params)
	}
}

func NewAppError(where string, id string, params map[string]interface{}, details string, status int) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = status
	ap.IsOAuth = false
	ap.Translate(translateFunc)
	return ap
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

// NewId is a globally unique identifier. It is a [A-Z0-9] string 26
// characters long. It is a UUID version Guid that is zbased32 encoded
// with the padding stripped off.
func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

// GetMillis is a convenience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
