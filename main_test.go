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

func TestWlog(t *testing.T) {
	//Info("Info testing start!!!")
	//Error("Error testing start!!!")
	//Debug("Debug testing start!!!")
	l := New()
	l.CallFrame()
	//l.SetLevel(PanicLevel)
	l.Json()
	l.Console()
	//l.Format = &JsonFormat{
	//	DisableTime:  true,
	//	DisableLevel: true,
	//	TimeFormat:   time.RFC850,
	//}
	f := make(Fields)
	f["key"] = "value"
	l.WithFields(f)
	l.WithField("field", "value")
	for i := 0; i < 10; i++ {
		//l.Debug("test Message")
		l.Info("test Message")
		//l.Warn("test Message")
		//l.Error("test Message")
		//l.Fatal("test Message")
		//l.Panic("test Message")

	}
}

func TestConsole(t *testing.T) {
	l := New()
	l.Console()
	log := New()
	log.Console()
	//l := New()
	l.Json()
	Info("test message111")
	l.Info("test message")
	//Panic("......")
	//f, err := os.Create("test.json")
	//if err != nil {
	//	return
	//}
	//defer f.Close()
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
	//l := New()
	l.Info("test message")
}

func BenchmarkNewLog(b *testing.B) {
	//l := New()
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
	//l := New()
	l.Json()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("test message")
		}
	})
}

func BenchmarkNewField(b *testing.B) {
	l := New()
	//l := New()
	f := make(map[string]any)
	f["field"] = "value"
	l.WithFields(f)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//l.WithFields(f).Info("test message")
			l.CallFramesDepth()
			l.Info("test message")
		}
	})
}

// func BenchmarkLogrusField(b *testing.B) {
// 	l := logrus.New()
// 	l.SetFormatter(new(logrus.JSONFormatter))
// 	f := make(logrus.Fields)
// 	f["out"] = "out"
// 	f["in"] = "in"
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			l.WithFields(f).Info("test message")
// 		}
// 	})
// }

func TestLog(t *testing.T) {
	l := New()
	//l := New()
	//l.IsCallFrame()
	//l.CallFramesDepth(1)
	for i := 0; i < 30; i++ {
		l.Infof("%s, %s, %s", "a", "b", "c")
		l.WithField(strconv.Itoa(i), i)
	}
}

func TestLogWithKeys(t *testing.T) {
	log := New()
	log.Console()
	log.WithKeys("a", "b", "c")
	for i := 1; i < 100; i++ {
		log.Values(i, i, i).Info("")
	}
}
func TestLogFile(t *testing.T) {
	log := New()
	log.Console()
	//log.Console()
	////log := New()
	for i := 0; i < 1000000; i++ {
		log.WithField("item", strconv.Itoa(i))
		log.Println(i)
	}

}
