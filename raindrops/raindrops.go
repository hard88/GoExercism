package raindrops

import (
	"strconv"
)

func Convert(drop int) string {
	answer := ""
	if drop % 3 == 0 {
		answer += "Pling"
	}

	if drop % 5 == 0 {
		answer += "Plang"
	}

	if drop % 7 == 0 {
		answer += "Plong"
	}

	if len(answer) == 0 {
		return strconv.Itoa(drop)
	} else {
		return answer
	}
}
