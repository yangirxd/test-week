package main

import "testing"

func TestIntersect(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, res := intersect(a, b)
	if !ok {
		t.Error("Expected intersection to exist")
	}
	want := []int{64, 3}
	if len(res) != len(want) {
		t.Fatalf("Expected %v, got %v", want, res)
	}
	for i := range want {
		if res[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, res)
		}
	}

	// Test no intersection
	a2 := []int{1, 2, 3}
	b2 := []int{4, 5, 6}
	ok2, res2 := intersect(a2, b2)
	if ok2 || len(res2) != 0 {
		t.Error("Expected no intersection")
	}
}
