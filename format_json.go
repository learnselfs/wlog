// Package wlog @Author Bing
// @Date 2024/2/22 16:32:00
// @Desc
package wlog

import "encoding/json"

const (
	Timestamp string = "time"
	LogLevel  string = "level"
	Message   string = "message"
	Errors    string = "error"
	Call      string = "call"
)

var (
	dataLength int
)

type JsonFormat struct{}

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
	data[LogLevel] = e.level
	data[Timestamp] = e.time
	data[Message] = e.msg
	data[Errors] = e.error
	if e.log.reportCaller {
		// todo: report call
	}
	for k, v := range e.data {
		data[k] = v
	}

	return data
}
