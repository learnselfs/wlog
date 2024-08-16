// Package wlog @Author Bing
// @Date 2024/8/12 16:11:00
// @Desc
package wlog

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestFileDestination(t *testing.T) {
	f := NewFileInfo()
	pwd, _ := os.Getwd()
	paths := []string{"", "./", pwd}
	for i, v := range paths {
		f.Path = v
		path, err := f.fileDestinationPath()
		if err != nil {
			t.Error(err)
		}
		t.Log(i, path)
		if path != pwd {
			t.Fatal(i, pwd, path)
		}
	}

}

func TestFileCronTask(t *testing.T) {
	Cron.Start()
	f := NewFileInfo()
	f.Cycle = SecondCycle
	f.cron()
	for i := 0; i < 3; i++ {
		err := f.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			t.Log(err)
		}
		time.Sleep(time.Second * 1)
	}

	ds, err := os.ReadDir(f.Path)
	if err != nil {
		t.Error(err)
	}
	if len(ds) <= 3 {
		t.Fatal("generate file number less 3: ", ds)
	}
	Cron.Stop()

}
