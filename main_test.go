// Package wlog @Author Bing
// @Date 2024/2/2 15:44:00
// @Desc
package wlog

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

//	func TestWlog(t *testing.T) {
//		Info.Println("Info testing start!!!")
//		Error.Println("Error testing start!!!")
//		Debug.Println("Debug testing start!!!")
//	}
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
	l.Info("afjiaejefajiofaj")
}

func BenchmarkNewLog(b *testing.B) {
	l := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			l.Info("aaaaaaaaaaaaaa")
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
//			l.Info("aaaaaaaaaaaaaa")
//		}
//	})
//}

func BenchmarkNewTextLog(b *testing.B) {
	l := New()
	l.SetFormatText()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("aaaaaaaaaaaaaaaa")
		}
	})
}

func BenchmarkNewField(b *testing.B) {
	l := New()
	f := make(Fields)
	f["out"] = "out"
	f["in"] = "in"
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.WithFields(f).Info("aaaaaaaaaaaaaaaa")
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
//			l.WithFields(f).Info("aaaaaaaaaaaaaaaa")
//		}
//	})
//}
