package binary

func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		min1 := abs(i*i - x)
		min2 := abs((i+1)*(i+1) - x)
		if min1 <= min2 {
			return i
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
