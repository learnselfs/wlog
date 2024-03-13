// Package wlog @Author Bing
// @Date 2024/2/2 15:38:00
// @Desc
package wlog

import "fmt"

func init() {
	initPool()
	log = New()
	log.SetLevel(InfoLevel)
}

func Debug(msg string) {
	log.Debug(msg)
}
func Info(msg string) {
	log.Info(msg)
}
func Warn(msg string) {
	log.Warn(msg)
}
func Error(msg string) {
	log.Error(msg)
}
func Fatal(msg string) {
	log.Fatal(msg)
}

func Panic(msg string) {
	log.Panic(msg)
}

func Debugf(f string, msg ...any) {
	log.ReportCaller()
	log.Debug(fmt.Sprintf(f, msg...))
}

func Infof(f string, msg ...any) {
	log.Info(fmt.Sprintf(f, msg...))
}
func Warnf(f string, msg ...any) {
	log.Warn(fmt.Sprintf(f, msg...))
}
func Errorf(f string, msg ...any) {
	log.Error(fmt.Sprintf(f, msg...))
}
func Fatalf(f string, msg ...any) {
	log.Fatal(fmt.Sprintf(f, msg...))
}

func Panicf(f string, msg ...any) {
	log.Panic(fmt.Sprintf(f, msg...))
}
func Print(msg string) {
	log.Print(msg)
}

func Println(msg string) {
	log.Println(msg)
}

func Printf(format, msg string) {
	log.Printf(format, msg)
}
