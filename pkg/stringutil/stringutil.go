package stringutil

import "unicode"

func ContainsUpper(str string) bool {

	for _, char := range str {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false

}
