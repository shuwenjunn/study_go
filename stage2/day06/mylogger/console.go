package mylogger

import (
	"fmt"
	"time"
)

//ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return c.Level <= LogLevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

// Debug 往终端写日志
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	if c.enable(DEBUG) {
		c.log(DEBUG, format, a...)
	}
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	if c.enable(INFO) {
		c.log(INFO, format, a...)
	}
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	if c.enable(WARNING) {
		c.log(WARNING, format, a...)
	}
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	if c.enable(ERROR) {
		c.log(ERROR, format, a...)
	}
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	if c.enable(FATAL) {
		c.log(FATAL, format, a...)
	}
}
