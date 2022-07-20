package _0findSubstring

import "reflect"

// 朴素哈希表
// 令n为字符串s的长度，m为数组words的长度（单词的个数），w为单词的长度。
// 由于words里面每个单词长度固定，而我们要找的字符串只能恰好包含所有的单词，因此我们要找到目标子串的长度为 m x w
// 思路：
// 1. 使用哈希表map记录words中每个单词的出现次数
// 2. 枚举s中的每个字符作为起点，往后取得长度为mxw的子串sub
// 3. 使用哈希表cur 统计sub每个单词的出现次数（每隔w长度作为一个单词）
// 4. 比较cur 和 map 是否相同
func findSubstring(s string, words []string) (res []int) {

	if len(words) == 0 {
		return
	}
	n, m, w := len(s), len(words), len(words[0])
	if m*w > n {
		return
	}

	wm := make(map[string]int, 0)
	for _, i := range words {
		wm[i]++
	}
	for i := 0; i+m*w <= n; i++ {
		cur := make(map[string]int, 0)
		for j := i; j < i+m*w; j += w {
			item := s[j : j+w]
			if _, ok := wm[item]; !ok {
				break
			}
			cur[item]++
		}

		// 性能是否可以优化
		if reflect.DeepEqual(wm, cur) {
			res = append(res, i)
		}
	}
	return
}
