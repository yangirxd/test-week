package main

import "testing"

func TestPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	go pipeline(in, out)
	go func() {
		in <- 2
		in <- 3
		close(in)
	}()
	results := []float64{}
	for v := range out {
		results = append(results, v)
	}
	if len(results) != 2 || results[0] != 8 || results[1] != 27 {
		t.Errorf("Expected [8 27], got %v", results)
	}
}
