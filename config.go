// Package wlog @Author Bing
// @Date 2024/2/21 16:18:00
// @Desc
package wlog

import (
	"bytes"
)

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	// format

)

var (
	//Debug *log.Logger
	//Info  *log.Logger
	//Warn  *log.Logger
	//Error *log.Logger
	//Fatal *log.Logger

	bufferPool BufferPool
	entryPool  EntryPool
)

type (
	Level int

	// Fields entry data type
	Fields map[string]interface{}

	// BufferPool interface
	BufferPool interface {
		Get() *bytes.Buffer
		Set(*bytes.Buffer)
	}

	// EntryPool interface
	EntryPool interface {
		Get() *Entry
		Set(*Entry)
	}

	// ReportFormat interface
	ReportFormat interface {
		Format(*Entry) ([]byte, error)
	}
)
