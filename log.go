// Package wlog @Author Bing
// @Date 2024/2/21 16:34:00
// @Desc
package wlog

import (
	"fmt"
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
	Format ReportFormat
	output io.Writer
	mu     *Mu
	fields Fields
}

// New create default logger
func New() *Log {
	l := &Log{
		level:        InfoLevel,
		reportCaller: false,
		bufferPool:   bufferPool,
		Format:       DefaultTextFormat(),
		output:       os.Stdout,
		mu:           NewMutex(),
		fields:       make(Fields),
	}
	NewEntryPool(l)
	l.entryPool = entryPool
	return l
}

// SetLevel  define log level
func (l *Log) SetLevel(level Level) {
	l.level = level
}

// SetJsonFormat define log output format
func (l *Log) SetJsonFormat() {
	l.Format = DefaultJsonFormat()
}

func (l *Log) isLevelEnabled(level Level) bool {
	return l.level <= level
}

func (l *Log) Debug(msg string) {
	entry := l.newEntry()
	entry.Debug(msg)
}
func (l *Log) Debugf(format string, msg ...any) {
	l.Debug(fmt.Sprintf(format, msg...))
}
func (l *Log) Info(msg string) {
	entry := l.newEntry()
	entry.Info(msg)
}

func (l *Log) Infof(format string, msg ...any) {
	l.Info(fmt.Sprintf(format, msg...))
}
func (l *Log) Warn(msg string) {
	entry := l.newEntry()
	entry.Warn(msg)
}
func (l *Log) Warnf(format string, msg ...any) {
	l.Warn(fmt.Sprintf(format, msg...))
}
func (l *Log) Error(msg string) {
	entry := l.newEntry()
	entry.Error(msg)
}
func (l *Log) Errorf(format string, msg ...any) {
	l.Error(fmt.Sprintf(format, msg...))
}
func (l *Log) Fatal(msg string) {
	entry := l.newEntry()
	entry.Fatal(msg)
}
func (l *Log) Fatalf(format string, msg ...any) {
	l.Fatal(fmt.Sprintf(format, msg...))
}

func (l *Log) Fatalln(msg string) {
	entry := l.newEntry()
	entry.Fatal(msg)
}

func (l *Log) Panic(msg string) {
	entry := l.newEntry()
	entry.Panic(msg)
}
func (l *Log) Panicln(msg string) {
	entry := l.newEntry()
	entry.Panic(msg)
}

func (l *Log) Panicf(format string, msg ...any) {
	entry := l.newEntry()
	entry.Panic(fmt.Sprintf(format, msg...))
}

func (l *Log) Print(msg string) {
	l.Info(msg)
}

func (l *Log) Println(msg string) {
	l.Info(msg)
}

func (l *Log) Printf(format string, msg ...any) {
	l.Info(fmt.Sprintf(format, msg...))
}

func (l *Log) newEntry() *Entry {
	entry := l.entryPool.Get()
	return entry
}

func (l *Log) releaseEntry(e *Entry) {
	l.entryPool.Set(e)
}

// WithFields appends fields to log
func (l *Log) WithFields(fields Fields) {
	for k, v := range fields {
		l.fields[k] = v
	}
}

func (l *Log) WithField(key string, value any) {
	f := make(Fields)
	f[key] = value
	l.WithFields(f)
}

// SetOutput define log output
func (l *Log) SetOutput(output io.Writer) {
	l.mu.Lock()
	defer l.mu.UnLock()
	l.output = output
}

// SetFormatter custom log formatter
func (l *Log) SetFormatter(f ReportFormat) {
	l.Format = f
}
