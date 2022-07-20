package _38findAnagrms

func findAnagrams2(s string, p string) (res []int) {

	ls := len(s)
	lp := len(p)
	if ls < lp {
		return
	}
	differ := make(map[byte]int, 0)
	for i := 0; i < lp; i++ {
		differ[p[i]]++
	}
	if len(differ) == 0 {
		res = append(res, 0)
	}

	for i := 0; i < ls; i++ {
		if i >= lp {
			differ[s[i-lp]]++
			// 加完之后，可能为0
			if differ[s[i-lp]] == 0 {
				delete(differ, s[i-lp])
			}
		}
		differ[s[i]]--
		// 减去之后可能为0
		if differ[s[i]] == 0 {
			delete(differ, s[i])
		}
		if len(differ) == 0 {
			res = append(res, i-lp+1)
		}
	}
	return res
}
