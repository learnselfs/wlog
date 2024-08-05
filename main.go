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

func Print(msg string) {
	log.Info(msg)
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
	log.Debug(fmt.Sprintln(msg...))
}
func Println(msg ...any) {
	log.Infoln(msg...)
}



func Infoln(msg ...any) {
	log.Infoln(msg...)
}
func Warnln(msg ...any) {
	log.Warn(fmt.Sprintln(msg...))
}
func Errorln(msg ...any) {
	log.Error(fmt.Sprintln(msg...))
}
func Fatalln(msg ...any) {
	log.Fatalln(msg...)
}

func Panicln(msg ...any) {
	log.Panicln(msg...)
}

func Debugf(f string, msg ...any) {
	log.Debugf(f, msg...)
}

func Printf(f string, msg ...any) {
	log.Infof(f, msg...)
}
func Infof(f string, msg ...any) {
	log.Infof(f, msg...)
}

func Warnf(f string, msg ...any) {
	log.Warnf(f, msg...)
}

func Errorf(f string, msg ...any) {
	log.Errorf(f, msg...)
}

func Fatalf(f string, msg ...any) {
	log.Fatalf(f, msg...)
}

func Panicf(f string, msg ...any) {
	log.Panicf(f, msg...)
}



func IsCallFrame() {
	log.IsCallFrame()
}

func CallFramesDepth(depths ...int) {
	log.CallFramesDepth(depths...)
}

func Json(){
	log.Json()
}

<<<<<<< HEAD
func Text(){
	log.Text()
}

=======
func ReportCaller() {
	log.ReportCaller()
}
>>>>>>> c6b78b30c3abf1111135f1998fcfc087cce61008
