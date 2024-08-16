// Package wlog @Author Bing
// @Date 2024/2/21 18:29:00
// @Desc
package wlog

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	log    *Log
	time   time.Time
	fields Fields
	msg    string
	frame  *runtime.Frame
	level  Level
	error  error
}

func NewEntry(log *Log) *Entry {
	return &Entry{
		log:    log,
		time:   time.Now(),
		fields: make(Fields),
	}
}

func (e *Entry) Dup() *Entry {
	fields := make(Fields, len(e.fields))
	for i, v := range e.fields {
		fields[i] = v
	}
	return &Entry{log: e.log, time: time.Now(), fields: fields, msg: e.msg}
}

func (e *Entry) handleLog(level Level, msg string) {
	//newEntry := e.Dup()
	//newEntry := e.log.newEntry()
	defer e.log.releaseEntry(e)
	e.fields = e.log.fields
	e.level = level
	e.msg = msg

	// handle entry data Fields
	// and handle format data
	e.write()

}

func (e *Entry) Log(level Level, msg string) {
	if e.log.Level(level) {
		e.handleLog(level, msg)
	} else {
		l, _ := level.Marshal()
		lvl, _ := e.log.level.Marshal()
		err := fmt.Sprintf("this log level is %s, not call %s()", lvl, l)
		e.error = errors.New(err)
		e.handleLog(ErrorLevel, "")
	}
}

func (e *Entry) Debug(msg string) {
	e.Log(DebugLevel, msg)
}
func (e *Entry) Info(msg string) {
	e.Log(InfoLevel, msg)
}
func (e *Entry) Warn(msg string) {
	e.Log(WarnLevel, msg)
}
func (e *Entry) Error(msg string) {
	e.Log(ErrorLevel, msg)
}
func (e *Entry) Fatal(msg string) {
	e.Log(FatalLevel, msg)
	os.Exit(1)
}
func (e *Entry) Panic(msg string) {
	e.Log(PanicLevel, msg)
	panic(msg)
}

func (e *Entry) withFields(fields Fields) {
	for k, v := range fields {
		e.fields[k] = v
	}
}

func (e *Entry) write() {
	if e.log.callFrame {
		e.frame = e.CallFrame(e.log.callFrameDepth)
	}
	byteData, err := e.log.Format.Format(e)
	if err != nil {
		return
	}
	if e.log.isOutput {
		_, err = e.log.output.Write(byteData)
	} else {
		err = e.log.File.Write(byteData)
		if err != nil {
			return
		}
	}
}
func (e *Entry) clear() {
	e.fields = make(Fields)
	e.msg = ""
}

func (e *Entry) CallFrame(depth int) *runtime.Frame {
	ptr := make([]uintptr, 24)
	runtime.Callers(0, ptr)
	var minPtr int
	// 将 caller 中包含 wlog package 摘出来
	for i, v := range ptr {
		f := runtime.FuncForPC(v)
		if strings.Contains(f.Name(), "wlog") {
			minPtr = i
			break
		}
	}
	frames := runtime.CallersFrames(ptr[minPtr:])
	for {
		f, _ := frames.Next()
		// 将 frame 中不包含 wlog package return
		if !strings.Contains(f.Function, "wlog") {
			if depth <= 0 {
				return &f
			}
			depth--
		}
	}
}

// Cron  linux cron task express:
// example:
//
//	"* * * * * *"
func (e *Entry) Cron(express string) {

}
