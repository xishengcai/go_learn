package longestPalindrome

import "testing"

func TestPalindrome(t *testing.T) {
	println(longestPalindrome("abcba"))
	println(longestPalindrome("xbcba"))
	println(longestPalindrome("abc"))
	println(longestPalindrome("cbbd"))
}

func TestPalindromeCenter(t *testing.T) {
	println(palindromeCenter("abcba"))
	println(palindromeCenter("abccba"))
	println(palindromeCenter("xbcba"))
	println(palindromeCenter("abc"))
	println(palindromeCenter("cbbd"))
}

func TestPalindromeDynamicPlan(t *testing.T) {
	println(palindromeDynamicPlan("abcba"))
	println(palindromeDynamicPlan("abccba"))
	println(palindromeDynamicPlan("xbcba"))
	println(palindromeDynamicPlan("abc"))
	println(palindromeDynamicPlan("cbbd"))
}

func TestPalindromeInt(t *testing.T) {
	t.Log(isPalindrome(12321))
	t.Log(isPalindrome(123321))
	t.Log(isPalindrome(0))
	t.Log(isPalindrome(10))
}
