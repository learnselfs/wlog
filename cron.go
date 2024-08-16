// Package wlog @Author Bing
// @Date 2024/8/14 17:12:00
// @Desc
package wlog

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const ()

type Task interface{ Run() }
type TaskFunc func() error

func (f TaskFunc) Run() {
	if err := f(); err != nil {
		fmt.Println(err)
	}
}

type cron struct {
	tasks   []*cronTask
	mu      sync.Mutex
	ticker  *time.Ticker
	disable chan struct{}
}

func newCron() *cron {
	return &cron{tasks: make([]*cronTask, 0), mu: sync.Mutex{}, ticker: time.NewTicker(time.Second), disable: make(chan struct{})}
}

type cronTask struct {
	second, minute, hour, day, month, week []int
	taskFunc                               Task
}

func (c *cronTask) schedule(t time.Time) bool {
	return container(c.second, t.Second()) && container(c.minute, t.Minute()) && container(c.hour, t.Hour()) && container(c.day, t.Day()) && container(c.month, int(t.Month())) && container(c.week, int(t.Weekday()))
}
func (cron *cron) Start() {
	go cron.run()
}

func (cron *cron) Stop() {
	cron.disable <- struct{}{}
}
func (cron *cron) run() {

	for {
		select {
		case t := <-cron.ticker.C:
			for _, c := range cron.tasks {
				if c.schedule(t) {
					c.taskFunc.Run()

				}
			}
		case <-cron.disable:
			cron.ticker.Stop()
			close(cron.disable)
			return
		}
	}
}

func Parser(express string, handle func() error) (*cronTask, error) {
	fields := strings.Fields(express)
	if len(fields) == 0 {
		return nil, errors.New("express is nil")
	}
	second, err := parts(fields[0], 0, 59)
	minute, err := parts(fields[1], 0, 59)
	hour, err := parts(fields[2], 0, 23)
	day, err := parts(fields[3], 0, 31)
	month, err := parts(fields[4], 1, 12)
	week, err := parts(fields[5], 0, 6)
	if err != nil {
		return nil, err
	}
	task := &cronTask{second, minute, hour, day, month, week, TaskFunc(handle)}
	Cron.tasks = append(Cron.tasks, task)
	return task, nil

}

// parts parser cron spacial
// -
// ,
// /
func parts(express string, min, max int) ([]int, error) {
	var fields []int
	var err error
	if len(express) == 0 {
		return nil, errors.New("parser express is nil ")
	}
	if len(express) == 1 {
		if express == "*" {
			fields = sliceGenerate(min, max)
		} else {
			duration, err := strconv.Atoi(express)
			if err != nil {
				return nil, errors.New("parser express failed: " + express)
			}
			fields = []int{duration}
		}
	}

	if strings.Contains(express, "/") {
		expresses := strings.Split(express, "/")
		if len(expresses) != 2 {
			return nil, errors.New("parser express failed: " + express)
		}
		var base []int
		base, err = parts(expresses[0], min, max)
		if err != nil {
			return nil, err
		}

		step, err := strconv.Atoi(expresses[1])
		if err != nil {
			return nil, errors.New("parser express failed: " + express)
		}
		for _, v := range base {
			if v%step == 0 {
				fields = append(fields, v)
			}
		}
	} else if strings.Contains(express, "-") {
		expresses := strings.Split(express, "-")
		l, err := strconv.Atoi(expresses[0])
		if err != nil {
			return nil, errors.New("parser express failed: " + expresses[0])
		}
		r, err := strconv.Atoi(expresses[1])
		if err != nil {
			return nil, errors.New("parser express failed: " + expresses[0])
		}
		fields = sliceGenerate(l, r)
	}
	if strings.Contains(express, ",") {
		expresses := strings.Split(express, ",")
		if len(express) < 2 {
			return nil, errors.New("parser express failed: " + express)
		}
		for _, value := range expresses {
			v, err := strconv.Atoi(value)
			if err != nil {
				return nil, errors.New("parser express failed: " + value)
			}
			fields = append(fields, v)
		}

	}

	return fields, err
}

func sliceGenerate(min, max int) []int {
	var fields []int
	for i := min; i <= max; i++ {
		fields = append(fields, i)
	}
	return fields

}

func container(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
