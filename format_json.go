// Package wlog @Author Bing
// @Date 2024/2/22 16:32:00
// @Desc
package wlog

import (
	"encoding/json"
	"errors"
	"time"
)

type JsonFormat struct {
	TimeFormat string // time format

	DisableTime  bool
	DisableLevel bool
	DisableColor bool
}

func (j *JsonFormat) Format(entry *Entry) ([]byte, error) {

	entry.log.mu.Lock()
	defer entry.log.mu.UnLock()
	buffer := entry.log.bufferPool.Get()
	defer entry.log.bufferPool.Set(buffer)

	data := j.Parse(entry)
	d := json.NewEncoder(buffer)
	err := d.Encode(&data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (j *JsonFormat) Parse(e *Entry) Fields {

	fields := make(Fields)
	var level string
	var err error
	if !j.DisableColor {
		level, err = e.level.MarshalColor()
	} else {
		level, err = e.level.Marshal()
	}
	if err != nil {
		e.error = errors.Join(e.error, err)
	}
	if e.error != nil {
		fields[Errors] = e.error.Error()
	}
	if !j.DisableLevel {
		fields[LogLevel] = level
	}
	if !j.DisableTime {
		fields[Timestamp] = e.time.Format(j.TimeFormat)
	}
	if e.msg == "" {
	} else {
		fields[Message] = e.msg
	}
	//data[Errors] = e.error
	if e.log.callFrame {
		fields[CallFile] = e.frame.File
		fields[CallLine] = e.frame.Line
		fields[CallFunc] = e.frame.Function
	}
	for k, v := range e.fields {
		fields[k] = v
	}

	return fields
}

func NewFormatJson() *JsonFormat {
	return &JsonFormat{
		TimeFormat:   time.DateTime,
		DisableColor: true,
	}
}
