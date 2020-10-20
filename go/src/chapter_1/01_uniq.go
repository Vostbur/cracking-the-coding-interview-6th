// Реализуйте алгоритм, определяющий, все ли символы в строке встречаются только один раз.
// А если при этом запрещено использование дополнительных структур данных?
package chapter_1

func IsUnique(s string) bool {
	d := make(map[rune]struct{})
	for _, i := range s {
		if _, ok := d[i]; ok {
			return false
		}
		d[i] = struct{}{}
	}
	return true
}
