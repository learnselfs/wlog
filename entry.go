// Package wlog @Author Bing
// @Date 2024/2/21 18:29:00
// @Desc
package wlog

import (
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
		data: make(Fields, 6),
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
	newEntry := e.log.entryPool.Get()
	defer e.log.entryPool.Set(newEntry)
	newEntry.level = level
	newEntry.msg = msg

	// handle entry data Fields
	// and handle format data
	newEntry.write()
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
}

func (e *Entry) withFields(fields Fields) {
	data := e.data
	for i, v := range fields {
		data[i] = v
	}
}

func (e *Entry) write() {
	byteData, err := e.log.format.Format(e)
	if err != nil {
		return
	}
	_, err = e.log.output.Write(byteData)
	if err != nil {
		return
	}
}
