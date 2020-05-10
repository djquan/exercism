package matrix

import (
	"errors"
	"strconv"
	"strings"
)

//Matrix represents a mathematical matrix
type Matrix [][]int

//New creates a new matrix from a given string
func New(s string) (Matrix, error) {
	var columnLength int
	var err error
	rows := strings.Split(s, "\n")
	result := make([][]int, len(rows))

	for i, line := range rows {
		elements := strings.Split(strings.TrimSpace(line), " ")
		result[i] = make([]int, len(elements))

		if columnLength == 0 {
			columnLength = len(elements)
		} else if columnLength != len(elements) {
			return nil, errors.New("invalid matrix")
		}

		for j, elem := range elements {
			if result[i][j], err = strconv.Atoi(elem); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

//Rows retrieves the rows of a Matrix
func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m))

	for i, row := range m {
		newRow := make([]int, len(row))
		copy(newRow, row)
		result[i] = newRow
	}

	return result
}

//Cols retrieves the columns of a Matrix
func (m Matrix) Cols() [][]int {
	result := make([][]int, len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		result[i] = make([]int, len(m))
	}

	for i, row := range m {
		for j, elem := range row {
			result[j][i] = elem
		}
	}

	return result
}

//Set sets the value of an element within a matrix
func (m Matrix) Set(r int, c int, i int) bool {
	if r < 0 || c < 0 {
		return false
	}

	if r >= len(m) || c >= len(m[0]) {
		return false
	}

	m[r][c] = i
	return true
}
