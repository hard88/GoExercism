package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(str string) bool {
	var alphabet = map[rune]int{}

	for _, r := range strings.ToLower(str){
		if ! unicode.IsLetter(r) {
			continue
		}
		if alphabet[r] == 0 {
			alphabet[r] = 1
		} else {
			return false
		}
	}
	return true
}
