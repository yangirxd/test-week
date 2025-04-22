package main

import (
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	ch := make(chan int)
	stop := make(chan struct{})
	go randomGenerator(ch, stop)
	for i := 0; i < 3; i++ {
		v := <-ch
		if v < 0 || v >= 100 {
			t.Errorf("Value out of range: %d", v)
		}
	}
	close(stop)
}
