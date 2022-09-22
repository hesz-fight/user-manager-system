package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// return the log level info
func (l *Level) String() string {
	switch *l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warm"
	case LevelError:
		return "error"
	case LevelFatal:
		return "Fatal"
	case LevelPanic:
		return "panic"
	}

	return ""
}

type Fields map[string]interface{}

type Logger struct {
	loglogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NerLogger(w io.Writer, prifix string, flag int) *Logger {
	l := log.New(w, prifix, flag)

	return &Logger{loglogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l

	return &nl
}

// 返回新的 Logger 并且修改对应的字段
func (l *Logger) WithContext(ctx context.Context) *Logger {
	nl := l.clone()
	nl.ctx = ctx

	return nl
}

func (l *Logger) WithCaller(skip int) *Logger {
	nl := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		nl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return nl
}

func (l *Logger) WithField(fields Fields) *Logger {
	nl := l.clone()
	if nl.fields == nil {
		nl.fields = make(Fields)
	}

	for k, v := range fields {
		nl.fields[k] = v
	}

	return nl
}

// 日志输出方法
func (l *Logger) output(level Level, message string) {

	body, _ := json.Marshal(l.JSONFormat(level, message))
	context := string(body)

	switch level {
	case LevelDebug, LevelInfo, LevelWarn, LevelError:
		l.loglogger.Print(context)
	case LevelFatal:
		l.loglogger.Fatal(context)
	case LevelPanic:
		l.loglogger.Panic(context)
	}
}

func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {

	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers

	if len(l.fields) > 0 { // 判空
		for k, v := range l.fields {
			if _, ok := data[k]; !ok { // 不存在
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WithCaller(1).output(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithCaller(1).output(LevelInfo, fmt.Sprintf(format, v...))
}
