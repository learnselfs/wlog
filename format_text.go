// Package wlog @Author Bing
// @Date 2024/2/25 19:26:00
// @Desc
package wlog

import (
	"fmt"
	"sort"
	"time"
)

type TextFormat struct {
	// TimeFormat
	TimeFormat string
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
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(fmt.Sprintf("\"%v\"", data[k]))
		b.WriteString("\t")
	}
	b.WriteString("\n")
	return b.Bytes(), nil

}

func (t *TextFormat) parse(e *Entry) Fields {

	if e.log.reportCaller {
		dataLength = 4
	}
	dataLength = 3

	data := make(Fields, dataLength+len(e.data))
	level, err := e.level.Marshal()
	if err != nil {
		data[" "+Errors] = err
	}
	data[" "+LogLevel] = level
	data[" "+Timestamp] = e.time.Format(t.TimeFormat)
	data[Message] = e.msg

	for k, v := range data {
		data[k] = v
	}
	return data
}

func DefaultTextFormat() *TextFormat {
	return &TextFormat{
		TimeFormat: time.DateTime,
	}
}
