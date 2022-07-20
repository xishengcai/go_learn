package _38findAnagrms

import "testing"

func TestFindAnagroms(t *testing.T) {

	t.Log(findAnagrams("cbaebabacd", "abc"))
	t.Log(findAnagrams("abab", "ab"))
	t.Log(findAnagrams("ababababab", "aab"))
}

func TestFindAnagroms2(t *testing.T) {

	t.Log(findAnagrams2("cbaebabacd", "abc"))
	t.Log(findAnagrams2("abab", "ab"))
	t.Log(findAnagrams2("ababababab", "aab"))
}
