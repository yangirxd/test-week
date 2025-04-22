package main

import "fmt"

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{m: make(map[string]int)}
}

func (s *StringIntMap) Add(key string, value int) {
	s.m[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.m, key)
}

func (s *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int, len(s.m))
	for k, v := range s.m {
		copyMap[k] = v
	}
	return copyMap
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.m[key]
	return v, ok
}

func main() {
	s := NewStringIntMap()
	s.Add("one", 1)
	s.Add("two", 2)
	fmt.Println("Exists 'one':", s.Exists("one"))
	s1, ok := s.Get("two")
	fmt.Printf("Get %d, bool %t\n", s1, ok)
	s.Remove("one")
	fmt.Println("Exists 'one' after remove:", s.Exists("one"))
	copyMap := s.Copy()
	fmt.Println("Copy length:", len(copyMap))
}
