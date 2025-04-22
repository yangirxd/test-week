package main

import (
	"fmt"
)

func pipeline(in <-chan uint8, out chan<- float64) {
	for v := range in {
		out <- float64(v) * float64(v) * float64(v)
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)
	go pipeline(in, out)
	go func() {
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
		close(in)
	}()
	for v := range out {
		fmt.Println(v)
	}
}
