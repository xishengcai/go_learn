package spiral_matrix

func spiralMatrix(n int) [][]int {
	/*
		i: i=0, j++, 1,2,3..n, j< n
		j: j=n, i++, i< n
		i: i=n, j--, j>=0
		j: j=0, i--, i>0

	*/
	count := 0
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	max := n
	for i := 0; i < n/2+1; i++ {
		line, column := i, i
		for ; column < max && matrix[line][column] == 0; column++ {
			matrix[line][column] = count
			count++
		}

		column--
		for line++; line < max && matrix[line][column] == 0; line++ {
			matrix[line][column] = count
			count++
		}

		line--
		for column--; column >= 0 && matrix[line][column] == 0; column-- {
			matrix[line][column] = count
			count++
		}

		column++
		for line--; line > 0 && matrix[line][column] == 0; line-- {
			matrix[line][column] = count
			count++
		}
		max--
	}

	for _, k := range matrix {
		for _, xx := range k {
			print(xx, " ")
		}
		println()
	}
	return matrix
}
