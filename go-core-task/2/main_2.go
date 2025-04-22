package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(slice []int) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, num int) []int {
	return append(slice, num)
}

func copySlice(slice []int) []int {
	copyS := make([]int, len(slice))
	copy(copyS, slice)
	return copyS
}

func removeElement(slice []int, idx int) []int {
	if idx < 0 || idx >= len(slice) {
		return slice
	}
	return append(slice[:idx], slice[idx+1:]...)
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rnd.Intn(100)
	}
	fmt.Println("Оригинальный слайс:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Четные числа:", evenSlice)

	added := addElements(originalSlice, 99)
	fmt.Println("После добавления элемента:", added)

	copied := copySlice(originalSlice)
	fmt.Println("Скопированный слайс:", copied)
	originalSlice[0] = 999
	fmt.Println("После изменения оригинального слайса, скопированный:", copied)

	removed := removeElement(originalSlice, 2)
	fmt.Println("После удаелния элемента (idx = 2):", removed)
}
