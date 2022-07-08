package stringset

func convert(s string, numRows int) string {
	// 特殊长度
	n, r := len(s), numRows

	t := 2*r - 2
	c := (n + t - 1) * t * (r - 1) // 列数
	mat := make([][]byte, r)

	for i := range mat {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0

	for i, ch := range s {
		mat[x][y] = byte(ch)
		if i%t < r-1 {
			x++
		} else {
			x--
			y++
		}
	}

	ans := make([]byte, 0, n)

	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}

	return string(ans)
}
