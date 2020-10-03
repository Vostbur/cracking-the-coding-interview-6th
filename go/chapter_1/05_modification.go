// Существуют три вида модифицирующих операций над строками:
// вставка символа, удаление символа и замена символа. Напишите
// функцию, которая проверяет, находятся ли две строки на расстоянии
// одной модификации (или нуля модификаций).
// Пример:
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false
package chapter_1

func CheckReplace(s1, s2 string) bool {
	flag := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if flag {
				return false
			} else {
				flag = true
			}
		}
	}
	return true
}

func CheckInsert(s1, s2 string) bool {
	var minStr, maxStr string
	if len(s1) < len(s2) {
		minStr, maxStr = s1, s2
	} else {
		minStr, maxStr = s2, s1
	}
	for i := 0; i < len(maxStr); i++ {
		if maxStr[:i]+maxStr[i+1:] == minStr {
			return true
		}
	}
	return false
}

func Modification(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) == len(s2) {
		return CheckReplace(s1, s2)
	}
	return CheckInsert(s1, s2)
}
