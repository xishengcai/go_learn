package main

type Trie struct {
	children map[rune]*Trie
	match    bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie, 0),
		match:    false,
	}
}

func NewTrie(w rune) *Trie {
	return &Trie{
		children: make(map[rune]*Trie, 0),
		match:    false,
	}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}
	for _, w := range word {
		if this.children[w] == nil {
			this.children[w] = NewTrie(w)
		}
		this = this.children[w]
	}
	this.match = true
}

func (this *Trie) Search(word string) bool {
	ok := false
	for _, w := range word {
		if this, ok = this.children[w]; !ok {
			return false
		}
	}

	return this.match
}

func (this *Trie) StartsWith(prefix string) bool {
	ok := false
	for _, w := range prefix {
		if this, ok = this.children[w]; !ok {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
