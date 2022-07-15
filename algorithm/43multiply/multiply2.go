package _3multiply

import "strconv"

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ans := 0
	m, n := len(num1), len(num2)

	// 被乘数
	for i := n - 1; i >= 0; i-- {
		x := int(num2[i]-'0') * myPow(10, n-1-i)
		add := 0
		// 乘数
		for j := m - 1; j >= 0; j-- {
			y := int(num1[j] - '0')
			add += x * y * myPow(10, m-1-j)
		}
		ans += add
	}
	ansStr := ""
	for ans != 0 {
		c := ans % 10
		ans = ans / 10
		ansStr = strconv.Itoa(c) + ansStr
	}
	return ansStr
}

func myPow(x, n int) int {
	if n == 0 || x == 1 {
		return 1
	}

	if n < 0 {
		x = 1.0 / x
		n = n * -1
	}

	if n == 1 {
		return x
	}

	if n%2 == 1 {
		half := myPow(x, n/2)
		return half * half * x
	}
	half := myPow(x, n/2)
	return half * half
}
