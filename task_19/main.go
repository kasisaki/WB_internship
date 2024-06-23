package main

import (
	"bufio"
	"fmt"
	"github.com/kasisaki/WB_internship/task_19/line_reverse"
	"os"
	"strings"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите текст: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			return
		}

		fmt.Println("Перевертыш1:", reverse.StrReverse(strings.TrimSpace(text))) // т.к. в тексте содержится перенос строки
	}
}
