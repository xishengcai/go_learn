package _3multiply

import "strconv"

func multiply3(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n, m := len(num1), len(num2)
	pos := make([]int, n+m)

	//`num1[i] * num2[j]` will be placed at indices `[i + j`, `i + j + 1]`

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			x1 := int(num1[i] - '0')
			x2 := int(num2[j] - '0')
			product := x1 * x2

			p1 := i + j
			p2 := i + j + 1
			sum := pos[p2] + product
			pos[p1] += sum / 10
			pos[p2] = sum % 10

		}
	}

	ansStr := ""
	for i := 0; i < n+m; i++ {
		if i == 0 && pos[i] == 0 {
			break
		}
		ansStr = ansStr + strconv.Itoa(pos[i])
	}

	return ansStr
}
