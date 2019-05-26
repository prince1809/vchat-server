package mlog

var globalLogger *Logger

type LogFunc func(string, ...Field)



var Debug LogFunc = defaultDebugLog
var Info LogFunc = defaultInfoLog
var Warn LogFunc = defaultWarnLog
var Error LogFunc = defaultErrorLog
var Critical LogFunc = defaultCriticalLog
