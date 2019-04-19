package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	step := 0
	if input <= 0{
		return step, errors.New("the input should be positive")
	}

	for i := input; i > 1; {
		if i %2 == 0 {
			i = i/2
		} else {
			i = i * 3 + 1
		}
		step += 1
	}

	return step, nil

}
