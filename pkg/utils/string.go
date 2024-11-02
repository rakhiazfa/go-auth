package utils

import "strings"

func LcFirst(str string) (result string) {
	for i, v := range str {
		if i == 0 {
			result += strings.ToLower(string(v))
		} else {
			result += string(v)
		}
	}

	return
}
