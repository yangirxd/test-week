package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { ch1 <- 1; close(ch1) }()
	go func() { ch2 <- 2; close(ch2) }()
	out := mergeChannels(ch1, ch2)
	results := []int{}
	for v := range out {
		results = append(results, v)
	}
	if len(results) != 2 || (results[0] != 1 && results[1] != 1) || (results[0] != 2 && results[1] != 2) {
		t.Errorf("Unexpected merge result: %v", results)
	}
}
