// Package wlog @Author Bing
// @Date 2024/2/2 15:38:00
// @Desc
package wlog

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "\033[0;34m [Info] \033[0m", log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(os.Stdout, "\033[0;31m [Error] \033[0m", log.Ldate|log.Ltime|log.Llongfile)
	Debug = log.New(os.Stdout, "\033[0;33m [Debug] \033[0m", log.Ldate|log.Ltime|log.Llongfile)

}
