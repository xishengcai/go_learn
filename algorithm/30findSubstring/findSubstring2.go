package _0findSubstring

func findSubstring2(s string, words []string) (ans []int) {
	ls, m, w := len(s), len(words), len(words[0])

	differ := map[string]int{}
	for _, word := range words {
		differ[word]++
	}

	for i := 0; i <= ls-m*w; i++ {
		for j := i; j < i+m*w; j += w {
			word := s[j : j+w]
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		if len(differ) == 0 {
			ans = append(ans, i)
		}

		// TODO: 需要优化时间
		for j := i; j < i+m*w; j += w {
			word := s[j : j+w]
			differ[word]++
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

	}

	return
}
