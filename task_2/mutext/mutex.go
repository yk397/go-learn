package mutext

import (
	"sync"
	"sync/atomic"
)

//题目一

type Conter struct {
	Count uint
	lock  sync.Mutex
}

func (conter *Conter) Increment() {
	defer conter.lock.Unlock()
	conter.lock.Lock()
	for i := 0; i < 1000; i++ {
		conter.Count += 1
	}
}

// 题目二
type Counter2 struct {
	Count atomic.Int32
}

func (conter *Counter2) Increment() {

	for i := 0; i < 1000; i++ {
		conter.Count.Add(1)
	}
}
