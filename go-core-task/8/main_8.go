package main

import (
	"fmt"
	"sync/atomic"
)

type MyWaitGroup struct {
	count int32
	ch    chan struct{}
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{ch: make(chan struct{})}
}

func (wg *MyWaitGroup) Add(counter int) {
	atomic.AddInt32(&wg.count, int32(counter))
}

func (wg *MyWaitGroup) Done() {
	if atomic.AddInt32(&wg.count, -1) == 0 {
		close(wg.ch)
	}
}

func (wg *MyWaitGroup) Wait() {
	<-wg.ch
}

func main() {
	wg := NewMyWaitGroup()
	wg.Add(2)
	go func() {
		fmt.Println("Goroutine 1 done")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine 2 done")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All done!")
}
