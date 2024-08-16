// Package wlog @Author Bing
// @Date 2024/8/12 15:16:00
// @Desc
package wlog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	jsonPostfix = "json"
	textPostfix = "text"
	logPostfix  = "log"
)

type Cycle int

func (c Cycle) string() string {
	switch c {
	case SecondCycle:
		return "* * * * * *"
	case MinuteCycle:
		return "0 * * * * *"
	case HourCycle:
		return "0 0 * * * *"
	case DayCycle:
		return "0 0 0 * * *"
	case MonthCycle:
		return "0 0 0 1 * *"
	case WeekCycle:
		return "0 0 0 1 * 0"
	}
	return ""
}

type File struct {
	Name     string
	Postfix  string
	fullName string
	// file destination path
	Path string
	// file size byte
	Max  float64
	Size float64
	// number of file backup
	Number int
	// cycle
	Cycle      Cycle
	Expression string
	lock       sync.Mutex
	// file
	file *os.File
}

func NewFile(name, path string, max, size float64, number int, cycle Cycle, expression string) *File {
	return &File{
		Name: name,
		Path: path,
		Size: size, Number: number,
		Max:        max,
		Cycle:      cycle,
		Expression: expression,
		lock:       sync.Mutex{},
	}
}

func NewFileTest() *File {
	return NewFileName("test", DayCycle)
}

func NewFileInfo() *File {
	return NewFileName("info", DayCycle)
}
func NewFileName(name string, cycle Cycle) *File {
	return NewFileCycle(name, cycle, "")
}
func NewFileCycle(name string, cycle Cycle, expression string) *File {
	return NewFile(name, "./log", 1024*1024*100, 0, 3, cycle, expression)
}

func (f *File) fileDestinationPath() (string, error) {
	var pwd string
	var err error
	if len(f.Path) == 0 {
		pwd, err = os.Getwd()
		if err != nil {
			return "", err
		}
		f.Path = pwd
		return pwd, err
	} else if ok := filepath.IsAbs(f.Path); ok {
		return f.Path, nil
	}
	pwd, err = os.Getwd()
	if err != nil {
		return "", err
	}
	f.Path = filepath.Join(pwd, f.Path)
	return f.Path, nil
}

func (f *File) fileNameList() []string {
	if len(f.Postfix) == 0 {
		f.Postfix = logPostfix
	}
	if len(f.Name) == 0 {
		return []string{"test", f.Postfix}
	}
	postfix := filepath.Ext(f.Name)
	f.Name = f.Name[:len(f.Name)-len(postfix)]
	return []string{f.Name, f.Postfix}
}

func (f *File) fileName() {
	fileName := f.fileNameList()
	f.fullName = strings.Join(fileName, ".")
}

func (f *File) filePath() string {
	f.fileName()
	return filepath.Join(f.Path, f.fullName)
}

func (f *File) fileNameBackup() string {
	date := time.Now().Format("2006-01-02 15_04_05")
	fileNames := f.fileNameList()
	fileNames = []string{fileNames[0], date, fileNames[1]}
	fileName := strings.Join(fileNames, ".")
	filePath := filepath.Join(f.Path, fileName)
	return filePath

}
func (f *File) Rotate() error {
	f.lock.Lock()
	defer f.lock.Unlock()
	var err error
	_, err = f.file.Stat()
	if err != nil {
		return nil
	}
	backupPath := f.fileNameBackup()
	filePath := f.filePath()
	err = f.file.Close()

	if err = os.Rename(filePath, backupPath); err != nil {
		return err
	}
	err = f.openOrCreateFile()
	//file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	//f.file = file
	//f.size = 0

	return nil
}

// Write log to file
func (f *File) Write(data []byte) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	dataSize := float64(len(data))

	if dataSize > f.Max {
		return errors.New("data size exceeds max size")
	}

	if f.file == nil {
		if err := f.openOrCreateFile(); err != nil {
			return err
		}
	}

	if dataSize+f.Size > f.Max {
		if err := f.Rotate(); err != nil {
			return err
		}
	}
	l, err := f.file.Write(data)
	if err != nil {
		return err
	}
	f.Size += float64(l)

	return nil
}

func (f *File) openOrCreateFile() error {
	path, err := f.fileDestinationPath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	fileName := f.filePath()
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	f.file = file
	return nil
}

// cron Parse expression
func (f *File) cron() {
	if len(f.Expression) == 0 {
		f.Expression = f.Cycle.string()
	}
	f.cycleTask(f.Expression)
}

// cycleTask Parse expression and cycle execute
func (f *File) cycleTask(expression string) {
	_, err := Parser(expression, f.Rotate)
	if err != nil {
		fmt.Sprintln(err)
	}
}
