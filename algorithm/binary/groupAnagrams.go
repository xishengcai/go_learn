package binary

import "sort"

func groupAnagrams(strs []string) (res [][]string) {
	m := make(map[string][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}
