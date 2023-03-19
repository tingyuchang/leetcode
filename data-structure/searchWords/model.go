package searchWords

import "fmt"

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	current := this

	for _, c := range word {
		idx := c - 'a'
		if current.children[idx] == nil {
			current.children[idx] = &WordDictionary{}
		}
		current = current.children[idx]
	}

	current.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return searchHelper(this, word, 0)
}

func searchHelper(dictionary *WordDictionary, word string, idx int) bool {
	if idx == len(word) {
		return dictionary.isEnd
	}

	c := word[idx]

	if c == '.' {
		for _, v := range dictionary.children {
			if v != nil && searchHelper(v, word, idx+1) {
				return true
			}
		}
		return false
	} else {
		c = c - 'a'
		fmt.Println(word, c, idx)
		fmt.Println(*dictionary)
		if dictionary.children[c] == nil {
			return false
		}
		return searchHelper(dictionary.children[c], word, idx+1)
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
