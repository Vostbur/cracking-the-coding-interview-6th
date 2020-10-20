// Допустим, что существует метод isSubstring, проверяющий,
// является ли одно слово подстрокой другого. Для двух строк s1 и s2 напишите код,
// который проверяет, получена ли строка s2 циклическим сдвигом s1,
// используя только один вызов метода isSubstring
// (пример: слово waterbottle получено циклическим сдвигом erbottlewat).
package chapter_1

import "strings"

func isSubstring(input, sub string) bool {
	return strings.Contains(input, sub)
}

func Shift(s1, s2 string) bool {
	return len(s1) == len(s2) && isSubstring(s1 + s1, s2)
}
