package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type matrix [][]int

func (m *matrix) Rows() [][]int {
	rows := make([][]int, 0, len(*m))
	for _, row := range *m {
		r := make([]int, len(row))
		copy(r, row) // the tow slices should be same length
		rows = append(rows, r)
	}
	return rows
}

func (m *matrix) Cols() [][]int {
	cols := make([][]int, 0, len(*m))
	l := len((*m)[0])
	for i := 0; i < l; i++ { //列
		col := make([]int, 0)
		for j := 0; j < len(*m); j++ { //行
			col = append(col, (*m)[j][i])
		}
		cols = append(cols, col)
	}
	return cols
}

func (m *matrix) Set(r, c, v int) bool {
	if r >= len(*m) || r < 0 || c >= len((*m)[0]) || c < 0 {
		return false
	}
	(*m)[r][c] = v
	return true
}

func New(in string) (*matrix, error) {
	var m matrix
	rows := strings.Split(in, "\n")
	var length int
	for _, row := range rows {
		values := strings.Fields(row)
		if len(values) < 1 {
			return nil, errors.New("rows have no value")
		}
		if length == 0 {
			length = len(values)
		} else if len(values) != length {
			return nil, errors.New("the length of rows should be same")
		}
		r := make([]int, 0, len(values))
		for _, value := range values {
			v, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			r = append(r, v)
		}
		m = append(m, r)
	}

	return &m, nil

}
