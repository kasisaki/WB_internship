package main

import (
	"fmt"
)

// SetBit устанавливает i-й бит в 1
func SetBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// ClearBit устанавливает i-й бит в 0
func ClearBit(n int64, i uint) int64 {
	return n &^ (1 << i)
}

func main() {
	var n int64 = 0           // Переменная, с которой будем работать
	var bitPosition uint = 10 // Позиция бита, которую будем изменять

	// Установка 3-го бита в 1
	n = SetBit(n, bitPosition)
	fmt.Printf("После установки %d-го бита в 1: %064b\n", bitPosition, n)
	fmt.Println(n)

	// Установка 3-го бита в 0
	n = ClearBit(n, bitPosition)
	fmt.Printf("После установки %d-го бита в 0: %064b\n", bitPosition, n)

}
