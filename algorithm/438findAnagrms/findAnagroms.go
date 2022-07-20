package _38findAnagrms

func findAnagrams(s string, p string) (res []int) {

	size := len(s)
	sizeP := len(p)

	if size < sizeP {
		return
	}

	//pCount := make([]int,26)
	//sCount := make([]int,26)
	// ? make 不可使用 == 比较

	var sCount, pCount [26]int
	for _, ch := range p {
		pCount[ch-'a']++
	}

	for i := 0; i < size; i++ {
		sCount[s[i]-'a']++
		if i >= sizeP {
			sCount[s[i-sizeP]-'a']--
		}
		if sCount == pCount {
			res = append(res, i+1-sizeP)
		}
	}
	return res
}
