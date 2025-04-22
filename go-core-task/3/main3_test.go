package main

import "testing"

func TestStringIntMap(t *testing.T) {
	s := NewStringIntMap()
	s.Add("a", 1)
	if !s.Exists("a") {
		t.Error("Expected key 'a' to exist")
	}
	v, ok := s.Get("a")
	if !ok || v != 1 {
		t.Errorf("Expected to get 1, got %v, %v", v, ok)
	}
	s.Remove("a")
	if s.Exists("a") {
		t.Error("Expected key 'a' to be removed")
	}
	s.Add("b", 2)
	copyMap := s.Copy()
	if len(copyMap) != 1 || copyMap["b"] != 2 {
		t.Error("Copy failed")
	}
	copyMap["b"] = 100
	v2, _ := s.Get("b")
	if v2 == 100 {
		t.Error("Copy is not independent")
	}
}
