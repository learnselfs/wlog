// Package wlog @Author Bing
// @Date 2024/2/27 15:57:00
// @Desc
package wlog

import "errors"

func (l Level) Marshal() (string, error) {
	switch l {
	case DebugLevel:
		return "debug", nil
	case InfoLevel:
		return "info", nil
	case WarnLevel:
		return "warn", nil
	case ErrorLevel:
		return "error", nil
	case FatalLevel:
		return "fatal", nil
	case PanicLevel:
		return "panic", nil
	default:
		return "", errors.New("not found level")

	}
}
