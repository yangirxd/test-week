package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// 1. Создание переменных
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	// 2. Определение типов
	fmt.Printf("numDecimal: %T\n", numDecimal)
	fmt.Printf("numOctal: %T\n", numOctal)
	fmt.Printf("numHexadecimal: %T\n", numHexadecimal)
	fmt.Printf("pi: %T\n", pi)
	fmt.Printf("name: %T\n", name)
	fmt.Printf("isActive: %T\n", isActive)
	fmt.Printf("complexNum: %T\n", complexNum)

	// 3. Преобразование в строку и объединение
	combined := combineToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Combined string:", combined)

	// 4. Преобразование строки в срез рун
	runes := []rune(combined)
	fmt.Println("Runes:", runes)

	// 5. Хэширование с солью
	hash := hashWithSalt(runes, "go-2024")
	fmt.Printf("SHA256 hash: %x\n", hash)
}

func combineToString(numDecimal, numOctal, numHexadecimal int, pi float64, name string, isActive bool, complexNum complex64) string {
	return fmt.Sprintf("%d|%d|%d|%f|%s|%t|%v", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
}

func hashWithSalt(runes []rune, salt string) [32]byte {
	mid := len(runes) / 2
	runesWithSalt := make([]rune, 0, len(runes)+len([]rune(salt)))
	runesWithSalt = append(runesWithSalt, runes[:mid]...)
	runesWithSalt = append(runesWithSalt, []rune(salt)...)
	runesWithSalt = append(runesWithSalt, runes[mid:]...)
	b := []byte(string(runesWithSalt))
	return sha256.Sum256(b)
}
