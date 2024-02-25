// Package wlog @Author Bing
// @Date 2024/2/22 15:13:00
// @Desc
package wlog

import (
	"bytes"
	"sync"
)

type DefaultPool struct {
	pool *sync.Pool
}

func (p *DefaultPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}
func (p *DefaultPool) Set(b *bytes.Buffer) {
	b.Reset()
	p.pool.Put(b)
}

type DefaultEntryPool struct {
	pool *sync.Pool
}

func (d *DefaultEntryPool) Get() *Entry {
	return d.pool.Get().(*Entry)
}
func (d *DefaultEntryPool) Set(e *Entry) {
	e.msg = ""
	e.data = nil
	d.pool.Put(e)
}
func init() {
	bufferPool = &DefaultPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}
func NewEntryPool(l *Log) {

	entryPool = &DefaultEntryPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return NewEntry(l)
			},
		},
	}
}
