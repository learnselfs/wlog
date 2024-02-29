// Package wlog @Author Bing
// @Date 2024/2/2 15:44:00
// @Desc
package wlog

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"testing"
)

func TestWlog(t *testing.T) {
	//Info("Info testing start!!!")
	//Error("Error testing start!!!")
	//Debug("Debug testing start!!!")
	l := New()
	l.SetJsonFormat()
	f := make(Fields)
	f["key"] = "value"
	l.WithFields(f)
	l.WithField("field", "value")
	for i := 0; i < 10; i++ {
		l.Info("test Message")
	}
}

func TestSetOutput(t *testing.T) {
	l := New()
	l.SetJsonFormat()
	l.Info("test message")
	Info("test message")
	//Panic("......")
	f, err := os.Create("test.json")
	if err != nil {
		return
	}
	defer f.Close()
	l.SetOutput(f)
	l.Info("test message")

}
func testRuntime() (uintptr, string, int, bool) {
	return runtime.Caller(1)
}

func testRuntimes() []string {
	var result []string
	slices := make([]uintptr, 16)
	runtime.Callers(0, slices)
	for _, v := range slices {
		if v != 0 {
			fun := runtime.FuncForPC(v)
			name := fun.Name()
			file, line := fun.FileLine(v)
			result = append(result, file+":"+name+":"+strconv.Itoa(line))
		}
	}
	return result
}

func TestRuntime(t *testing.T) {
	pc, file, line, ok := testRuntime()
	if ok {
		fn := runtime.FuncForPC(pc)
		fmt.Println(file, line, fn.Name(), fn.Entry())
	}
	var a int
	fmt.Println(a)
}

func TestRuntimes(t *testing.T) {
	slices := make([]uintptr, 100)
	runtime.Callers(0, slices)
	result := testRuntimes()
	fmt.Println(result)
	fmt.Println(slices)
}

func TestNewLog(t *testing.T) {
	l := New()
	l.Info("test message")
}

func BenchmarkNewLog(b *testing.B) {
	l := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			l.Info("test message")
		}
	})
}

//func BenchmarkLogrus(b *testing.B) {
//	l := logrus.New()
//	f := new(logrus.TextFormatter)
//	f.DisableSorting = true
//	l.SetFormatter(f)
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//
//			l.Info("test message")
//		}
//	})
//}

func BenchmarkNewTextLog(b *testing.B) {
	l := New()
	l.SetJsonFormat()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("test message")
		}
	})
}

func BenchmarkNewField(b *testing.B) {
	l := New()
	f := make(Fields)
	f["field"] = "value"
	l.WithFields(f)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//l.WithFields(f).Info("test message")
			l.Info("test message")
		}
	})
}

//func BenchmarkLogrusField(b *testing.B) {
//	l := logrus.New()
//	l.SetFormatter(new(logrus.JSONFormatter))
//	f := make(logrus.Fields)
//	f["out"] = "out"
//	f["in"] = "in"
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			l.WithFields(f).Info("test message")
//		}
//	})
//}
