// Напишите метод, заменяющий все пробелы в строке символами '%20'.
// Можете считать, что длина строки позволяет сохранить дополнительные символы,
// а фактическая длина строки известна заранее.
// Пример:
// Ввод: "Mr John Smith    ", 13
// Вывод: "Mr%20John%20Smith"
package chapter_1

func URLify(input string, length int) string {
	tmp := make([]rune, len(input))
	for i, count := 0, 0; i < length; i++ {
		if input[i] == ' ' {
			tmp[count] = rune('%')
			tmp[count+1] = rune('2')
			tmp[count+2] = rune('0')
			count += 2
		} else {
			tmp[count] = rune(input[i])
		}
		count++
	}
	return string(tmp)
}

func URLifyInPlace(input []rune, length int) {
	index := len(input) - 1
	for i := length - 1; i >= 0; i-- {
		if input[i] == rune(' ') {
			input[index-2] = rune('%')
			input[index-1] = rune('2')
			input[index] = rune('0')
			index -= 3
		} else {
			input[index] = input[i]
			index--
		}
	}
}
