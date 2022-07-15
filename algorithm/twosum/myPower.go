package twosum

func myPow(x float64, n int) float64 {

	if n == 0 || x == 1 {
		return 1
	}

	if n < 0 {
		x = 1.0 / x
		n = n * -1
	}

	res := float64(1)
	for ; n > 0; n-- {
		res = res * x
	}

	return res
}

func myPow2(x float64, n int) float64 {
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
		half := myPow2(x, n/2)
		return half * half * x
	}
	half := myPow2(x, n/2)
	return half * half

}
