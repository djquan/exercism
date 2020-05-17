package pascal

func Triangle(size int) [][]int {
	result := make([][]int, size)

	for line := 1; line <= size; line++ {
		l := make([]int, line)

		c := 1
		for i := 1; i <= line; i++ {
			l[i-1] = c
			c = c * (line - i) / i
		}

		result[line-1] = l
	}

	return result
}
