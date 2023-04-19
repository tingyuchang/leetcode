package AuthenticationManager

import (
	"container/heap"
)

// TODO using heap to resolve this
type TokenId struct {
	len     byte
	payload [5]byte
}

func NewTokenId(tokenId string) TokenId {
	t := TokenId{}
	t.len = byte(len(tokenId))
	copy(t.payload[:], tokenId)
	return t
}

type Token struct {
	id        TokenId
	ttl       int
	heapIndex int
}

type tokenHeap []*Token

func (t tokenHeap) Len() int {
	return len(t)
}

func (t tokenHeap) Less(i, j int) bool {
	return t[i].ttl < t[j].ttl
}

func (t tokenHeap) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].heapIndex = j
	t[j].heapIndex = i
}

func (t *tokenHeap) Push(x interface{}) {
	token := x.(*Token)
	token.heapIndex = len(*t)
	*t = append(*t, token)
}

func (t *tokenHeap) Pop() interface{} {
	token := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return token
}

type AuthenticationManager struct {
	ttl       int
	tokens    tokenHeap
	idToToken map[TokenId]*Token
}

func ConstructorAuthenticationManager(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		idToToken: make(map[TokenId]*Token),
		tokens:    make(tokenHeap, 0, 1),
		ttl:       timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	id := NewTokenId(tokenId)
	if token := this.idToToken[id]; token != nil {
		token.ttl = currentTime + this.ttl
		heap.Fix(&this.tokens, int(token.heapIndex))
	} else {
		token = &Token{
			id:  id,
			ttl: currentTime + this.ttl,
		}
		heap.Push(&this.tokens, token)
		this.idToToken[id] = token
	}
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if token := this.idToToken[NewTokenId(tokenId)]; token != nil && token.ttl > currentTime {
		token.ttl = currentTime + this.ttl
		heap.Fix(&this.tokens, int(token.heapIndex))
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	for len(this.tokens) != 0 && this.tokens[0].ttl <= currentTime {
		delete(this.idToToken, this.tokens[0].id)
		heap.Pop(&this.tokens)
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
