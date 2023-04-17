package Trie

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	root := &Trie{children: make(map[byte]*Trie)}

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		current := root

		for j := 0; j < len(strs[i]); j++ {
			_, ok := current.children[strs[i][j]]

			if !ok {
				current.children[strs[i][j]] = &Trie{children: make(map[byte]*Trie)}
			}

			if j == len(strs[i])-1 {
				current.isEnd = true
			}
			current = current.children[strs[i][j]]
		}
	}

	current := root

	prefix := ""
	for current != nil {

		if len(current.children) == 1 {
			var temp *Trie
			for k, v := range current.children {
				//fmt.Println(v, string(k))
				prefix += string(k)
				temp = v
			}
			if current.isEnd {
				break
			} else {
				current = temp
			}
		} else {
			break
		}
	}

	return prefix
}
