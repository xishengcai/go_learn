package _0findSubstring

// ls = len(s), lw = len(words), le = len(words[0]), lt = lw * le
// for i = 0 ~ le-1 作为开始位置开始一个滑动窗口，每个滑动窗口的结果合并在一起
// 对于每个滑动窗口, 开始滑动，窗口起始位置计为 start， end计为窗口的结束位置
// map[ele][count] 维护 words 中每个单词出现的次数， count 维护待匹配单词的个数
// 1. 每次初始时，end == start; 当 end开始，长度为le的单次在 map 中存在，且计数大于0，map[s] = map[s] - 1; 其中s是从end开始，长度为le的单词；窗口扩张，end = end + le
// 2. 当end开始，长度为le的单词在map中不存在(计数为0)，则窗口收缩， map[s] = map[s] + 1;其中s是从start开始，长度为le的单词 start = start + le
// 3. 当count == 0 时，意味着窗口扩张到目标大小 意味着已start开始的lt大小的窗口满足要求，start就是一个可行解
// 4. 当 start == end，窗口已经缩小为0
// 滑动窗口的起始位置应该 小于等于 ls - lt
// 当一个滑动窗口结束时，为了把map重置成原有的样子，需要从当前start位置收缩到end位置，把窗口内元素放到map中去
func findSubstring3(s string, words []string) []int {
	ls, lw, le := len(s), len(words), len(words[0])
	lt := lw * le
	if ls < lt {
		return []int{}
	}

	dict := make(map[string]int)
	for _, word := range words {
		dict[word] = dict[word] + 1
	}

	var res []int
	for i := 0; i < le; i++ {
		res = findInSlide(s, i, le, dict, lw, res)
	}
	return res
}

func findInSlide(s string, i int, step int, dict map[string]int, count int, res []int) []int {
	totalCount := count // 记录下来，防止改变
	start, end := i, i
	for ; start <= len(s)-step*totalCount; start += step { // 窗口收缩
		for {
			if count == 0 { // 窗口大小满足预期
				res = append(res, start)
				str := s[start : start+step]
				dict[str] = dict[str] + 1
				count++
				break
			}
			if end >= len(s) {
				break
			}
			str := s[end : end+step]
			if dict[str] > 0 { // 该字符有效，窗口可以继续扩张
				dict[str] = dict[str] - 1
				count--
				end += step
			} else { // 字符无效，窗口不能继续扩张，开始收缩
				if start == end { // 当前窗口已经为0， start = end = start + step // 窗口后移
					end += step
					break
				} else { // 窗口大小不为0，后窗口不动
					str := s[start : start+step]
					dict[str] = dict[str] + 1
					count++
					break
				}
			}
		}
	}
	for start < end {
		str := s[start : start+step]
		dict[str] = dict[str] + 1
		count++
		start += step
	}
	return res
}
