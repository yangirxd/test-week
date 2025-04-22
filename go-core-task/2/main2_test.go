package main

import "testing"

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}
	got := sliceExample(input)
	if len(got) != len(want) {
		t.Fatalf("Expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	got := addElements(input, 4)
	want := []int{1, 2, 3, 4}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	copyS := copySlice(input)
	if len(copyS) != len(input) {
		t.Errorf("Copy length mismatch")
	}
	copyS[0] = 100
	if input[0] == 100 {
		t.Errorf("Copy is not independent")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got := removeElement(input, 1)
	want := []int{1, 3, 4}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}
}
