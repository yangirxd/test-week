package main

import "testing"

func TestDiffSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	want := []string{"apple", "cherry", "43", "lead", "gno1"}
	got := diffSlice(slice1, slice2)
	if len(got) != len(want) {
		t.Fatalf("Expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}
}
