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
	entry := l.newEntry()
	entry.Debug(msg)
}
func (l *Log) Info(msg string) {
	entry := l.newEntry()
	entry.Info(msg)
}

func (l *Log) Warn(msg string) {
	entry := l.newEntry()
	entry.Warn(msg)
}
func (l *Log) Error(msg string) {
	entry := l.newEntry()
	entry.Error(msg)
}
func (l *Log) Fatal(msg string) {
	entry := l.newEntry()
	entry.Fatal(msg)
}

func (l *Log) newEntry() *Entry {
	entry := l.entryPool.Get()
	return entry
}

func (l *Log) releaseEntry(e *Entry) {
	l.entryPool.Set(e)
}

func (l *Log) WithFields(fields Fields) *Entry {
	entry := l.newEntry()
	entry.withFields(fields)
	return entry
}
