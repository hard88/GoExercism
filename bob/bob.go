// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.


// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isSlience(str string) bool {
	return len(str) == 0
}

func isShouting(str string) bool {
	hasLetters := strings.IndexFunc(str, unicode.IsLetter) >= 0
	isUpCase := strings.Compare(str, strings.ToUpper(str)) == 0
	return hasLetters && isUpCase
}

func isQuestion(str string) bool {
	return strings.HasSuffix(str, "?")
}

func isExasperated(str string) bool {
	return isShouting(str) && isQuestion(str)
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	var answer string
	newRemark := strings.TrimSpace(remark)

	switch {
	case isSlience(newRemark):
		answer = "Fine. Be that way!"
	case isExasperated(newRemark):
		answer = "Calm down, I know what I'm doing!"
	case isShouting(newRemark):
		answer = "Whoa, chill out!"
	case isQuestion(newRemark):
		answer = "Sure."
	default:
		answer = "Whatever."
	}

	return answer
}
