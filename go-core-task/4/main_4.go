package main

import "fmt"

func diffSlice(slice1, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice2))
	for _, v := range slice2 {
		m[v] = struct{}{}
	}
	result := make([]string, 0)
	for _, v := range slice1 {
		if _, found := m[v]; !found {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := diffSlice(slice1, slice2)
	fmt.Println("Result:", res)
}
