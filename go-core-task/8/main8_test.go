package main

import (
	"testing"
	"time"
)

func TestMyWaitGroup(t *testing.T) {
	wg := NewMyWaitGroup()
	wg.Add(2)
	flag := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Done()
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(flag)
	}()
	select {
	case <-flag:
		// ok
	case <-time.After(100 * time.Millisecond):
		t.Error("WaitGroup did not finish in time")
	}
}
