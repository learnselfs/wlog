// Package wlog @Author Bing
// @Date 2024/2/22 17:18:00
// @Desc
package wlog

import "sync"

type Mu struct {
	isEnable bool
	lock     sync.Mutex
}

func (m *Mu) Lock() {
	if !m.isEnable {
		m.lock.Lock()
	}
}

func (m *Mu) UnLock() {
	if !m.isEnable {
		m.lock.Unlock()
	}
}

func NewMutex() *Mu {
	return &Mu{
		isEnable: true,
		lock:     sync.Mutex{},
	}
}
