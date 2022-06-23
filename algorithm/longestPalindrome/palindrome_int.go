package longestPalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0

	for reverse < x {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return x == reverse || x == reverse/10
}
