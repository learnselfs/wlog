// Package wlog @Author Bing
// @Date 2024/8/14 23:19:00
// @Desc
package wlog

import (
	"fmt"
	"testing"
	"time"
)

func TestCronParser(t *testing.T) {
	specs := []string{
		"* * * * * *",
		"0 10-30/5 * * * *",
		"1-20 10-30/5 * * * *",
	}
	handle := func() error { fmt.Println("a"); return nil }
	for _, spec := range specs {
		if _, err := Parser(spec, handle); err != nil {
			t.Error(err)
		}
	}
}

func TestCronTaskRun(t *testing.T) {
	f1 := func() error { fmt.Println(1); return nil }
	f2 := func() error { fmt.Println(2); return nil }
	express := map[string]func() error{
		"* * * * * *":    f1,
		"0-59 * * * * *": f2,
	}
	Cron.Start()

	for k, v := range express {
		_, err := Parser(k, v)
		if err != nil {
			t.Error(err)
		}
	}
	time.Sleep(5 * time.Second)

	Cron.Stop()
}
