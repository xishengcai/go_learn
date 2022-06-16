package main

func main() {
	println(lengthOfLongestSubstring2("abcabcbb"))
	println(lengthOfLongestSubstring2("pwwkew"))
	println(lengthOfLongestSubstring2("bbbbb"))
	println(lengthOfLongestSubstring2(" "))
}

func lengthOfLongestSubstring(s string) int {
	count, max := 0, 0
	l := len(s)
	for i := 0; i < l; i++ {
		count = 1
		temp := map[byte]int{}
		temp[s[i]] = 1
		for j := i + 1; j < l; j++ {
			if temp[s[j]] == 1 {
				break
			}
			temp[s[j]]++
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	l := len(s)
	max, j := 0, 1
	temp := map[byte]int{}
	for i := 0; i+1 < l; i++ {
		for j < l && temp[s[j]] == 0 {
			temp[s[j]]++
			j++
		}
		count := j - i - 1
		if count > max {
			max = count
		}
		delete(temp, s[i+1])
	}
	return max
}
