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
	// entry pool
	entryPool  EntryPool  // log instance pool
	bufferPool BufferPool // log buffer pool

	// format
	Format ReportFormat // log output format : json/text

	isOutput bool // log output stream: file/stdout
	output   io.Writer
	mu       *Mu
	fields   Fields   // log fields
	keys     []string // log fields keys

	// report print caller information
	callFrame      bool // log call frame enable : false
	callFrameDepth int  // log call frame depth: default 0

	// file
	*File
	// cron
	*cron
}

func New() *Log {
	return NewLogInfoJson()
}
func NewLogInfoJson() *Log {
	return NewLog(InfoLevel, NewFileInfo(), NewFormatJson())
}

// NewLog create info logger
func NewLog(level Level, file *File, format ReportFormat) *Log {
	return NewLogConfig(level, false, 0, false, os.Stdout, file, format)
}

// NewLogConfig create default logger
func NewLogConfig(level Level, callFrame bool, callFrameDepth int, isOutput bool, output io.Writer, file *File, format ReportFormat) *Log {
	l := &Log{
		level:          level,
		callFrame:      callFrame,
		bufferPool:     bufferPool,
		Format:         format,
		isOutput:       isOutput,
		output:         output,
		mu:             NewMutex(),
		fields:         make(Fields),
		callFrameDepth: callFrameDepth,
		File:           file,
		cron:           Cron,
	}
	NewEntryPool(l)
	l.File.cron()
	//l.entryPool = entryPool
	return l
}

// Cron run task
//func (l *Log) Cron(expression string) {
//	l.File.cron(expression)
//}

// SetLevel  define log level
func (l *Log) SetLevel(level Level) {
	l.level = level
}

func (l *Log) CallFramesDepth(depths ...int) {
	l.callFrame = true
	if len(depths) == 0 {
		l.callFrameDepth = 1
	} else {
		l.callFrameDepth = depths[0]
	}
}

// Json SetJsonFormat define log output format
func (l *Log) Json() {
	l.Format = NewFormatJson()
}

// Text SetJsonFormat define log output format
func (l *Log) Text() {
	l.Format = NewFormatText()
}

func (l *Log) Level(level Level) bool {
	return l.level <= level
}

func (l *Log) Debug(msg string) {
	entry := l.newEntry()
	entry.Debug(msg)
}

func (l *Log) Debugln(msg ...any) {
	l.Debug(fmt.Sprintln(msg...))
}

func (l *Log) Debugf(format string, msg ...any) {
	l.Debug(fmt.Sprintf(format, msg...))
}

func (l *Log) Info(msg string) {
	entry := l.newEntry()
	entry.Info(msg)
}

func (l *Log) Infoln(msg ...any) {
	l.Info(fmt.Sprintln(msg...))
}

func (l *Log) Infof(format string, msg ...any) {
	l.Info(fmt.Sprintf(format, msg...))
}

func (l *Log) Warn(msg string) {
	entry := l.newEntry()
	entry.Warn(msg)
}

func (l *Log) Warnln(msg ...any) {
	l.Warn(fmt.Sprintln(msg...))
}

func (l *Log) Warnf(format string, msg ...any) {
	l.Warn(fmt.Sprintf(format, msg...))
}

func (l *Log) Error(msg string) {
	entry := l.newEntry()
	entry.Error(msg)
}

func (l *Log) Errorln(msg ...any) {
	l.Error(fmt.Sprintln(msg...))
}

func (l *Log) Errorf(format string, msg ...any) {
	l.Error(fmt.Sprintf(format, msg...))
}

func (l *Log) Fatal(msg string) {
	entry := l.newEntry()
	entry.Fatal(msg)
}

func (l *Log) Fatalln(msg ...any) {
	l.Fatal(fmt.Sprintln(msg...))
}

func (l *Log) Fatalf(format string, msg ...any) {
	l.Fatal(fmt.Sprintf(format, msg...))
}

func (l *Log) Panic(msg string) {
	entry := l.newEntry()
	entry.Panic(msg)
}

func (l *Log) Panicln(msg ...any) {
	l.Panic(fmt.Sprintln(msg...))
}

func (l *Log) Panicf(format string, msg ...any) {
	l.Panic(fmt.Sprintf(format, msg...))
}

func (l *Log) Print(msg string) {
	l.Info(msg)
}

func (l *Log) Println(msg ...any) {
	l.Print(fmt.Sprintln(msg...))
}

func (l *Log) Printf(f string, msg ...any) {
	l.Print(fmt.Sprintf(f, msg...))
}

func (l *Log) newEntry() *Entry {
	entry := l.entryPool.Get()
	return entry
}

func (l *Log) releaseEntry(e *Entry) {
	l.entryPool.Set(e)
}

// WithFields appends fields to log
func (l *Log) WithFields(fields map[string]any) {
	//l.fields = make(Fields)
	//for k, v := range fields {
	//	l.fields[k] = v
	//}
	l.fields = fields
}

func (l *Log) WithField(key string, value any) {
	l.WithFields(map[string]any{key: value})
}

func (l *Log) WithKeys(keys ...string) {
	l.keys = keys
}
func (l *Log) Values(values ...any) *Log {
	fields := make(map[string]any)
	for i, k := range l.keys {
		fields[k] = values[i]
	}
	l.WithFields(fields)
	return l
	//entry := l.newEntry()
	//entry.Log(l.level, "")

}

// Console enable default console output log
func (l *Log) Console() {
	l.isOutput = !l.isOutput
}

func (l *Log) Output(output io.Writer) {
	l.output = output
}

// Formatter custom log formatter
func (l *Log) Formatter(f ReportFormat) {
	l.Format = f
}

func (l *Log) JsonFormatDetail(timeFormat string, disableTime, disableColor, disableLevel bool) {
	l.Format = &JsonFormat{TimeFormat: timeFormat, DisableTime: disableTime, DisableColor: disableColor, DisableLevel: disableLevel}
}

func (l *Log) JsonTime(timeFormat string) {
	l.JsonFormatDetail(timeFormat, false, false, false)
}

func (l *Log) JsonTimeDisable() {
	l.JsonFormatDetail("", true, false, false)
}

func (l *Log) JsonColorDisable(timeFormat string) {
	l.JsonFormatDetail(timeFormat, false, true, false)
}

func (l *Log) TextFormatDetail(timeFormat string, disableTime, disableColor, disableLevel bool) {
	l.Format = &TextFormat{TimeFormat: timeFormat, DisableTime: disableTime, DisableColor: disableColor, DisableLevel: disableLevel}
}

func (l *Log) TextTime(timeFormat string) {
	l.TextFormatDetail(timeFormat, false, false, false)
}

func (l *Log) TextTimeDisable() {
	l.TextFormatDetail("", true, false, false)
}

func (l *Log) TextColorDisable(timeFormat string) {
	l.TextFormatDetail(timeFormat, false, true, false)
}

func (l *Log) CallFrame() {
	l.callFrame = !l.callFrame
}
