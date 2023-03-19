package PrefixTrie

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, ch := range word {
		idx := ch - 'a'
		if current.children[idx] == nil {
			current.children[idx] = &Trie{}
		}
		current = current.children[idx]
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, ch := range word {
		idx := ch - 'a'
		if current.children[idx] == nil {
			return false
		}
		current = current.children[idx]
	}
	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, ch := range prefix {
		idx := ch - 'a'
		if current.children[idx] == nil {
			return false
		}
		current = current.children[idx]
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
