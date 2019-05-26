package mlog

import (
	"encoding/json"
	"fmt"
)

// defaultLog manually encodes the log to STDOUT, providing a basic, default logging implementation
// before mlog is fully configured.
func defaultLog(level, msg string, fields ...Field) {
	log := struct {
		Level   string  `json:"level"`
		Message string  `json:"msg"`
		Fields  []Field `json:"fields,omitempty"`
	}{
		Level:   level,
		Message: msg,
		Fields:  fields,
	}

	if b, err := json.Marshal(log); err != nil {
		fmt.Printf(`{"level": "error", "msg":"failed to encode log message"}%s`, "\n")
	} else {
		fmt.Printf("%s\n", b)
	}
}

func defaultErrorLog(msg string, fields ...Field) {
	defaultLog("error", msg, fields...)
}
