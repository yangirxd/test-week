package main

import (
	"fmt"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)
	count := len(channels)
	for _, ch := range channels {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			count--
			if count == 0 {
				close(out)
			}
		}(ch)
	}
	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		close(ch2)
	}()

	go func() {
		ch3 <- 3
		close(ch3)
	}()

	for v := range mergeChannels(ch1, ch2, ch3) {
		fmt.Println(v)
	}
}
