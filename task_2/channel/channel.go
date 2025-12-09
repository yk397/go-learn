package channel

import (
	"fmt"
	"sync"
	"time"
)

// 题目一
func Communication() {

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("recived and print:", <-ch)
		}
	}()

	wg.Wait()

}

func BufferChan() {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("recived and print:", <-ch)
		}
	}()

	wg.Wait()
}
