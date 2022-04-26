package tree

type BTree struct {
	Value int
	R *BTree
	L *BTree
}

