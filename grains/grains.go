package grains

import (
	"errors"
)

func Square(squareNumber int) (uint64, error) {
	if squareNumber < 1 || squareNumber > 64 {
		return 0, errors.New("invalid")
	}

	return uint64(1 << uint(squareNumber -1)), nil
}

func Total() uint64 {
	return uint64((1 << 64) -1)
}

