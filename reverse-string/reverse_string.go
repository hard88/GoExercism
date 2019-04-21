package reverse

func reverse(str []rune) string {
	for i, j := 0, len(str)-1;  i < j; i, j = i + 1, j - 1 {
		str[i], str[j] = str[j], str[i]
	}

	return string(str)
}

func String(str string) string {
	return reverse([]rune(str))
}
