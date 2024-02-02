// Package wlog @Author Bing
// @Date 2024/2/21 16:34:00
// @Desc
package wlog

import (
	"io"
	"os"
)

type Log struct {
	level Level

	// report print caller information
	reportCaller bool
	// entry pool
	entryPool  EntryPool
	bufferPool BufferPool
	// format
	format ReportFormat
	output io.Writer
	mu     *Mu
}

func New() *Log {
	l := &Log{
		level:        InfoLevel,
		reportCaller: false,
		bufferPool:   bufferPool,
		format:       new(JsonFormat),
		output:       os.Stdout,
		mu:           NewMutex(),
	}
	NewEntryPool(l)
	l.entryPool = entryPool
	return l
}

func (l *Log) SetLevel(level Level) {
	l.level = level
}

func (l *Log) SetFormatText() {
	l.format = new(defaultTextFormat)
}

func (l *Log) isLevelEnabled(level Level) bool {
	return l.level >= level
}

func (l *Log) Debug(msg string) {
	entry := NewEntry(l)
	entry.Debug(msg)
}
func (l *Log) Info(msg string) {
	entry := NewEntry(l)
	entry.Info(msg)
}

func (l *Log) Warn(msg string) {
	entry := NewEntry(l)
	entry.Warn(msg)
}
func (l *Log) Error(msg string) {
	entry := NewEntry(l)
	entry.Error(msg)
}
func (l *Log) Fatal(msg string) {
	entry := NewEntry(l)
	entry.Fatal(msg)
}

func (l *Log) newEntry() *Entry {
	entry := l.entryPool.Get()
	return entry
}

func (l *Log) releaseEntry(e *Entry) {
	e.data = make(Fields)
	l.entryPool.Set(e)
}

func (l *Log) WithFields(fields Fields) *Entry {
	entry := l.newEntry()
	defer l.releaseEntry(entry)
	entry.withFields(fields)
	return entry
}
