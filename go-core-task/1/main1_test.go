package main

import (
	"fmt"
	"testing"
)

func TestCombineToString(t *testing.T) {
	res := combineToString(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	expected := "42|42|42|3.140000|Golang|true|(1+2i)"
	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestHashWithSalt(t *testing.T) {
	input := []rune("teststring")
	salt := "go-2024"
	hash := hashWithSalt(input, salt)
	// Ожидаемый хэш для строки с солью, вставленной в середину
	expectedHash := "436223bc069bb1e957e03febc83cfb15acf58a266b5a7ee30220ab2ffcc33af4"
	if actual := fmt.Sprintf("%x", hash); actual != expectedHash {
		t.Errorf("Expected %s, got %s", expectedHash, actual)
	}
}
