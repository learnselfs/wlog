// Package wlog @Author Bing
// @Date 2024/2/25 19:26:00
// @Desc
package wlog

import (
	"bytes"
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

	data := t.parse(e)
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k == Message {
			continue
		}
		writeKeyValue(b, k, data[k])
	}
	writeKeyValue(b, Message, data[Message])
	b.WriteString("\n")
	return b.Bytes(), nil

}

func (t *TextFormat) parse(e *Entry) Fields {
	data := make(Fields)
	var level string
	var err error
	if !t.DisableColor {
		level, err = e.level.MarshalColor()
	} else {
		level, err = e.level.Marshal()
	}
	if err != nil {
		data[" "+Errors] = err
	}
	if !t.DisableLevel {
		data[" "+LogLevel] = level
	}
	if !t.DisableTime {
		data[" "+Timestamp] = e.time.Format(t.TimeFormat)
	}
	data[Message] = e.msg

	for k, v := range e.data {
		data[k] = v
	}
	return data
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
