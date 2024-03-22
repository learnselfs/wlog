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

func ReportCal() {
	log.ReportCaller()
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

func Debugln(msg ...any) {
	log.ReportCaller()
	log.Debug(fmt.Sprintln(msg...))
}

func Infoln(msg ...any) {
	log.Info(fmt.Sprintln(msg...))
}
func Warnln(msg ...any) {
	log.Warn(fmt.Sprintln(msg...))
}
func Errorln(msg ...any) {
	log.Error(fmt.Sprintln(msg...))
}
func Fatalln(msg ...any) {
	log.Fatal(fmt.Sprintln(msg...))
}

func Panicln(msg ...any) {
	log.Panic(fmt.Sprintln(msg...))
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

func Print() {
	log.Print()
}

func Println() {
	log.Println()
}

func Printf() {
	log.Printf()
}
