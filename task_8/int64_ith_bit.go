package main

import (
	"fmt"
)

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// SetBit устанавливает i-й бит в 1
func SetBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// ClearBit устанавливает i-й бит в 0
func ClearBit(n int64, i uint) int64 {
	return n &^ (1 << i)
}

func main() {
	var n int64 = 512         // Переменная, с которой будем работать
	var bitPosition uint = 10 // Позиция бита, которую будем изменять

	// Установка 10-го бита в 1
	n = SetBit(n, bitPosition)
	fmt.Printf("После установки %d-го бита в 1: %064b\n", bitPosition, n)
	fmt.Println(n) // ожидаем 9й и 10й бит = 1, то есть 1024 + 512 = 1536

	// Установка 10-го бита в 0
	n = ClearBit(n, bitPosition)
	fmt.Printf("После установки %d-го бита в 0: %064b\n", bitPosition, n)
	fmt.Println(n) // ожидаем 9й бит равный 1, а 10 = 0, то есть 512

}
