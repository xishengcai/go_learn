package _0findSubstring

import "testing"

func TestFindSubstring(t *testing.T) {
	t.Log(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	t.Log(findSubstring2("barfoothefoobarman", []string{"foo", "bar"}))
	t.Log(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}
