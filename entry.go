// Package wlog @Author Bing
// @Date 2024/2/21 18:29:00
// @Desc
package wlog

import (
	"os"
	"time"
)

type Entry struct {
	log   *Log
	time  time.Time
	data  Fields
	msg   string
	level Level
	error error
}

func NewEntry(log *Log) *Entry {
	return &Entry{
		log:  log,
		time: time.Now(),
		data: make(Fields),
	}
}

func (e *Entry) Dup() *Entry {
	data := make(Fields, len(e.data))
	for i, v := range e.data {
		data[i] = v
	}
	return &Entry{log: e.log, time: time.Now(), data: data, msg: e.msg}
}

func (e *Entry) handleLog(level Level, msg string) {
	//newEntry := e.Dup()
	//newEntry := e.log.newEntry()
	defer e.log.releaseEntry(e)
	e.data = e.log.fields
	e.level = level
	e.msg = msg

	// handle entry data Fields
	// and handle format data
	e.write()

}

func (e *Entry) Log(level Level, msg string) {
	if e.log.isLevelEnabled(level) {
		e.handleLog(level, msg)
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
		e.data[k] = v
	}
}

func (e *Entry) write() {
	byteData, err := e.log.Format.Format(e)
	if err != nil {
		return
	}
	_, err = e.log.output.Write(byteData)
	if err != nil {
		return
	}
}

func (e *Entry) clear() {
	e.data = make(Fields)
	e.msg = ""
}
