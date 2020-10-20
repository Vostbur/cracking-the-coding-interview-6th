// Для двух строк напишите метод, определяющий, является ли одна строка перестановкой другой
package chapter_1

func Permutation(input string, output string) bool {
	if input == output || len(input) != len(output) {
		return false
	}

	d := make(map[rune]byte)
	for _, r := range input {
		d[r]++
	}

	for _, r := range output {
		if _, ok := d[r]; !ok {
			return false
		}
		if d[r] == 0 {
			return false
		} else {
			d[r]--
		}
	}
	return true
}
