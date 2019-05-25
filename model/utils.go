package model

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
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
