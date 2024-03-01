// Package wlog @Author Bing
// @Date 2024/2/25 19:26:00
// @Desc
package wlog

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"time"
)

type TextFormat struct {
	// TimeFormat
	TimeFormat string
	// Disable
	DisableTime  bool
	DisableLevel bool
	DisableColor bool
}

func (t *TextFormat) Format(e *Entry) ([]byte, error) {
	e.log.mu.Lock()
	defer e.log.mu.UnLock()
	b := e.log.bufferPool.Get()
	defer e.log.bufferPool.Set(b)

	fields := t.parse(e)
	var keys []string
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k == Message || k == Errors {
			continue
		}
		writeKeyValue(b, k, fields[k])
	}
	if fields[Message] == nil || fields[Message] == "" || fields[Errors] == nil {
	} else {
		writeKeyValue(b, Errors, fields[Errors])
		writeKeyValue(b, Message, fields[Message])
	}
	b.WriteString("\n")
	return b.Bytes(), nil

}

func (t *TextFormat) parse(e *Entry) Fields {
	fields := make(Fields)
	var level string
	var err error
	if !t.DisableColor {
		level, err = e.level.MarshalColor()
	} else {
		level, err = e.level.Marshal()
	}
	if err != nil {
		e.error = errors.Join(e.error, err)
	}
	if e.error != nil {
		fields[Errors] = e.error
	}
	if !t.DisableLevel {
		fields[" "+LogLevel] = level
	}
	if !t.DisableTime {
		fields[" "+Timestamp] = e.time.Format(t.TimeFormat)
	}
	if e.log.reportCaller {
		fields[CallFile] = e.frame.File
		fields[CallLine] = e.frame.Line
		fields[CallFunc] = e.frame.Function
	}
	fields[Message] = e.msg

	for k, v := range e.fields {
		fields[k] = v
	}
	return fields
}

func DefaultTextFormat() *TextFormat {
	return &TextFormat{
		TimeFormat: time.DateTime,
	}
}

func writeKeyValue(b *bytes.Buffer, key string, value interface{}) {
	b.WriteString(key)
	b.WriteString("=")
	b.WriteString(fmt.Sprintf("\"%v\"", value))
	b.WriteString("\t")
}
