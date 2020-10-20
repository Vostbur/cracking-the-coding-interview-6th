// Напишие функцию, которая проверяет, является ли заданная строка перестановкой палиндрома.
// (Палиндром - слово или фраза, одинаково читающиеся в прямом и обратном направлении;
// перестановка - строка, содержащая те же символы в другом порядке.)
// Палиндром не ограничивается словами из словаря.
// Пример: Tact Coa
// Вывод: True (перестановки: 'taco cat', 'atco cta' и т.д.)
package chapter_1

import (
	"log"
	"regexp"
	"strings"
)

func Palindrome(input string) bool {
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	tmp := reg.ReplaceAllString(strings.ToLower(input), "")
	l := len(tmp)

	counter := make(map[rune]int)
	for _, r := range tmp {
		counter[r]++
	}

	odd := l%2 == 1
	flag := false
	for _, value := range counter {
		if value%2 == 1 {
			if flag {
				return false
			}
			flag = true
		}
	}
	return flag == odd
}
