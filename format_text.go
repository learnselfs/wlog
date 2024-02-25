// Package wlog @Author Bing
// @Date 2024/2/25 19:26:00
// @Desc
package wlog

import (
	"fmt"
	"sort"
)

type defaultTextFormat struct {
}

func (f *defaultTextFormat) Format(e *Entry) ([]byte, error) {
	e.log.mu.Lock()
	defer e.log.mu.UnLock()
	b := e.log.bufferPool.Get()
	defer e.log.bufferPool.Set(b)

	//j := json.NewEncoder(b)
	data := f.parse(e)
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(fmt.Sprintf("%v", data[k]))
		b.WriteString("\t")
	}
	b.WriteString("\n")
	return b.Bytes(), nil

}

func (f *defaultTextFormat) parse(e *Entry) Fields {

	if e.log.reportCaller {
		dataLength = 4
	}
	dataLength = 3

	data := make(Fields, dataLength+len(e.data))
	data[LogLevel] = e.level
	data[Timestamp] = e.time
	data[Message] = e.msg

	for k, v := range data {
		data[k] = v
	}
	return data
}
