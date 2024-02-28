// Package wlog @Author Bing
// @Date 2024/2/22 16:32:00
// @Desc
package wlog

import (
	"encoding/json"
	"time"
)

const ()

var (
	dataLength int
)

type JsonFormat struct {
	TimeFormat string // time format
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
	if e.log.reportCaller {
		dataLength = 4
	} else {
		dataLength = 3
	}

	data := make(Fields, len(e.data)+dataLength)
	level, err := e.level.Marshal()
	if err != nil {
		data[Errors] = err.Error()
	}
	data[LogLevel] = level
	data[Timestamp] = e.time.Format(j.TimeFormat)
	data[Message] = e.msg
	//data[Errors] = e.error
	if e.log.reportCaller {
		// todo: report call
	}
	for k, v := range e.data {
		data[k] = v
	}

	return data
}

func DefaultJsonFormat() *JsonFormat {
	return &JsonFormat{
		TimeFormat: time.DateTime,
	}
}
