package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(ch chan int, stop chan struct{}) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		select {
		case <-stop:
			close(ch)
			return
		case ch <- rnd.Intn(100):
		}
	}
}

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go randomGenerator(ch, stop)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	close(stop)
}
