package mylogger

//往终端写日志相关内容

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type loglevel uint16

const (
	UNKNOWN loglevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(levelStr string) (loglevel, error) {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getlogString(lv loglevel) (string, error) {
	switch lv {
	case DEBUG:
		return "DEBUG", nil
	case TRACE:
		return "TRACE", nil
	case INFO:
		return "INFO", nil
	case WARNING:
		return "WARNING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FETAL", nil
	default:
		err := errors.New("无效的日志级别")
		return "UNKNOWN", err
	}
}

func getInfo(skip int) (fileName string, funcName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file)
	return
}
