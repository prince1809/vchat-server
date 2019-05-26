package mlog

var globalLogger *Logger

type LogFunc func(string, ...Field)




var Error LogFunc = defaultErrorLog
