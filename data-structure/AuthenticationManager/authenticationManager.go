package AuthenticationManager

import "fmt"

// TODO using heap to resolve this

type Token struct {
	id        string
	ttl       int
	heapIndex int
}

type TokenHeap []*Token

func (t *TokenHeap) minHeapButtonUp(n int) {
	current := n
	for current > 0 {
		parrent := (current - 1) / 2
		if (*t)[current].ttl < (*t)[parrent].ttl {
			(*t)[current], (*t)[parrent] = (*t)[parrent], (*t)[current]
			(*t)[current].heapIndex = current
			(*t)[parrent].heapIndex = parrent
		} else {
			break
		}
		current = parrent
	}
}

func (t *TokenHeap) minHeapTopDown(n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1

	var smallest int

	if l < len(*t) && (*t)[l].ttl < (*t)[n].ttl {
		smallest = l
	} else {
		smallest = n
	}

	if r < len(*t) && (*t)[r].ttl < (*t)[smallest].ttl {
		smallest = r
	}

	if smallest != n {
		(*t)[n], (*t)[smallest] = (*t)[smallest], (*t)[n]
		(*t)[n].heapIndex = n
		(*t)[smallest].heapIndex = smallest
		t.minHeapTopDown(smallest)
	}

}

func (t *TokenHeap) removeElement(n int) {
	// swap index n & last element
	(*t)[n], (*t)[len(*t)-1] = (*t)[len(*t)-1], (*t)[n]
	(*t)[n].heapIndex = n
	(*t)[len(*t)-1].heapIndex = len(*t) - 1
	*t = (*t)[:len(*t)-1]
	t.minHeapTopDown(n)

}

type AuthenticationManager struct {
	ttl       int
	tokens    TokenHeap
	idToToken map[string]*Token
}

func ConstructorAuthenticationManager(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl:       timeToLive,
		tokens:    make([]*Token, 0),
		idToToken: make(map[string]*Token),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	if _, ok := this.idToToken[tokenId]; ok {
		return
	}

	token := Token{id: tokenId, ttl: currentTime + this.ttl}
	token.heapIndex = len(this.tokens)
	this.tokens = append(this.tokens, &token)
	this.tokens.minHeapButtonUp(token.heapIndex)
	this.idToToken[tokenId] = &token
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if _, ok := this.idToToken[tokenId]; !ok {
		return
	}
	token := this.idToToken[tokenId]
	if token.ttl <= currentTime {
		// drop token
		delete(this.idToToken, tokenId)
		// remove from heap
		this.tokens.removeElement(token.heapIndex)
	} else {
		token.ttl = currentTime + this.ttl
		// run minHeapTop Down
		this.tokens.minHeapTopDown(token.heapIndex)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {

	for i, v := range this.tokens {
		fmt.Printf("%d %v\n", i, *v)
	}

	if len(this.tokens) == 0 {
		return 0
	}

	token := this.tokens[0]

	for token.ttl <= currentTime {
		delete(this.idToToken, token.id)
		this.tokens.removeElement(0)

		if len(this.tokens) == 0 {
			break
		}
		token = this.tokens[0]
	}

	return len(this.tokens)
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
