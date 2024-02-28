// Package wlog @Author Bing
// @Date 2024/2/2 15:38:00
// @Desc
package wlog

func init() {
	initPool()
	log = New()
	log.SetLevel(FatalLevel)
}

func Debug(msg string) {
	log.Info(msg)
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
func Panicln(msg string) {
	log.Panic(msg)
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
