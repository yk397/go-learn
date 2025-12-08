package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 题目一
func NumberPrinter(isEvent bool) {
	for i := 1; i < 11; i++ {
		if isEvent && (i%2 == 0) {
			fmt.Println("偶数打印器打印:", i)
		}
		if !isEvent && (i%2 == 1) {
			fmt.Println("奇数打印器打印:", i)
		}
	}
}

// 题目二
type TaskCoordinator struct {
	Tasks []func()
	Time  []int64
}

func (tc *TaskCoordinator) AddTask(fc ...func()) {
	tc.Tasks = append(tc.Tasks, fc...)
	tc.Time = make([]int64, len(tc.Tasks))
}

func (tc *TaskCoordinator) StartAllTask() {
	var wg sync.WaitGroup
	for i := range tc.Tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("goroutin-%d started...\n", i)
			startTime := time.Now()
			tc.Tasks[i]()
			endTime := time.Now()
			tc.Time[i] = endTime.UnixMilli() - startTime.UnixMilli()
			fmt.Printf("Task{%d} finised,consume %d MillSesonds\n", i, tc.Time[i])
		}(&wg)
	}
	wg.Wait()
}
