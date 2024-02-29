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
func (l Level) MarshalColor() (string, error) {
	switch l {
	case DebugLevel:
		return "\x1b[1;32mdebug\x1b[0m", nil
	case InfoLevel:
		return "\x1b[1;34minfo\x1b[0m", nil
	case WarnLevel:
		return "\x1b[1;33mwarn\x1b[0m", nil
	case ErrorLevel:
		return "\x1b[1;35merror\x1b[0m", nil
	case FatalLevel:
		return "\x1b[1;31mfatal\x1b[0m", nil
	case PanicLevel:
		return "\x1b[1;37mpanic\x1b[0m", nil
	default:
		return "", errors.New("not found level")
	}
}
