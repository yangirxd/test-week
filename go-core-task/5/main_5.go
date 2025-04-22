package main

import "fmt"

func intersect(a, b []int) (bool, []int) {
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	result := make([]int, 0)
	found := false
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
			found = true
		}
	}
	return found, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, res := intersect(a, b)
	fmt.Println("Intersection exists:", ok)
	fmt.Println("Intersection values:", res)
}
