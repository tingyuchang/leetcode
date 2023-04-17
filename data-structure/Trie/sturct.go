package Trie

type Trie struct {
	children map[byte]*Trie
	isEnd    bool
}
