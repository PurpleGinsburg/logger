package mylogger

import (
	"fmt"
	"time"
)

//往终端写日志相关内容

// 造一个loglevel定义日志级别

// logger对象
type Logger struct {
	level loglevel
}

// 构造函数
func Newlog(levelStr string) Logger {
	newlevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: newlevel,
	}
}

func (l Logger) enable(loglevel loglevel) bool {
	return loglevel >= l.level
}

func log(lv loglevel, msg string) {
	now := time.Now()
	fileName, funcName, lineNo := getInfo(3)
	lvstring, _ := getlogString(lv)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lvstring, fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}

}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		log(TRACE, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
