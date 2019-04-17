package hamming

import "errors"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) { return 0, errors.New("error") }

	if len(a) == 0 { return 0, nil }

	distance := 0
	for i:=0; i<len(a); i++{
		if a[i] != b[i] { distance++ }
	}
	return distance, nil
}
