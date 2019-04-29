package robotname

import (
	"errors"
	"math/rand"
	"strconv"
)

type Robot struct {
	name string
}

const maxSize = 26 * 26 * 10 * 10 * 10

var nameMap = make(map[string]int, maxSize)

func (r *Robot) Name() (string, error) {
	if len(nameMap) == maxSize {
		return "", errors.New("namespace is exhausted")
	}
	if len(r.name) == 0 {
		r.name = generateName()
	}
	return r.name, nil
}

func generateName() string {
	var name string
	for {
		name = ""
		name += string('A' + rune(rand.Intn(26)))
		name += string('A' + rune(rand.Intn(26)))
		name += strconv.Itoa(rand.Intn(10))
		name += strconv.Itoa(rand.Intn(10))
		name += strconv.Itoa(rand.Intn(10))
		if nameMap[name] < 1 {
			break
		}
	}
	nameMap[name] = 1
	return name
}

func (r *Robot) Reset() {
	r.name = ""
}
