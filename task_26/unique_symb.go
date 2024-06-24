package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:

	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func strUnique(s string) bool {
	m := make(map[rune]bool) // можно было использовать map[rune]struct{}, но так проверка выглядит лаконичнее

	s = strings.ToLower(s)

	for _, ch := range s {
		if m[ch] {
			return false //
		}
		m[ch] = true
	}
	return true
}

func main() {
	s1, s2, s3, s4, s5, s6 := "abcd", "acCdefAaf", "aabcd", "", "abc-d", "abc--d"

	fmt.Printf("%s — %v\n", s1, strUnique(s1)) // true
	fmt.Printf("%s — %v\n", s2, strUnique(s2)) // false
	fmt.Printf("%s — %v\n", s3, strUnique(s3)) // false
	fmt.Printf("%s — %v\n", s4, strUnique(s4)) // true
	fmt.Printf("%s — %v\n", s5, strUnique(s5)) // true
	fmt.Printf("%s — %v\n", s6, strUnique(s6)) // false
}
