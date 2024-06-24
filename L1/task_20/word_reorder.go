package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	s := "snow1 dog2 sun3 WB_intern4"
	rev := reorder(s)

	fmt.Printf("%s — %s\n", s, rev)
}

func reorder(s string) string {
	words := strings.Split(s, " ")
	var builder strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		builder.WriteRune(' ')
	}
	return strings.TrimSpace(builder.String())
}
