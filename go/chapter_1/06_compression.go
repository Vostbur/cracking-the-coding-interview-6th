// Реализуйте метод для выполнения простейшего сжатия строк с использованием
// счетчика повторяющихся символов. Например, строка aabcccccaa превращается в a2b1c5a3.
// Если 'сжатая' строка не становится короче исходной, то метод возвращает исходную строку.
// Предполагается, что строка состоит только из букв верхнего и нижнего регистра (a-z)
package chapter_1

import "strconv"

func Compression(input string) string {
	l := len(input)
	if l < 3 {
		return input
	}

	var result string
	count := 1
	for i := 1; i < l; i++ {
		if input[i] != input[i-1] {
			result = result + string(input[i-1]) + strconv.Itoa(count)
			count = 0
		}
		count++
	}
	result = result + string(input[l-1]) + strconv.Itoa(count)

	if len(result) < l {
		return result
	}
	return input
}

func CompressionWithSlice(input string) string {
	l := len(input)
	if l < 3 {
		return input
	}

	var r []byte
	count := 1
	for i := 1; i < l; i++ {
		if input[i] != input[i-1] {
			r = append(r, input[i-1])
			r = append(r, []byte(strconv.Itoa(count))...)
			count = 0
		}
		count++
	}
	r = append(r, input[l-1])
	r = append(r, []byte(strconv.Itoa(count))...)

	result := string(r)
	if l <= len(result) {
		return input
	}
	return result
}
