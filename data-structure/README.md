# Data structure

## Tree

### Full Binary Tree

- 所有internal node都有兩個subtree(也就是兩個child pointer)
- 所有leaf node具有相同的level(或相同的height)。

### Complete Binary Tree
若一棵樹的node按照Full Binary Tree的次序排列(由上至下，由左至右)，則稱此樹為Complete Binary Tree

## Heap

- Complete Binary Tree
- follow heap condition (max or min)
- 雖說 heap 在觀念上是一棵 complete binary tree, 實際上是存在一個陣列當中 -- root 存在 A[1], 接下來將 A[2] 與 A[3] 由左到右依序補滿第二層, 再將 A[4], A[5], A[6], A[7] 由左到右依序補滿第三層, 這種儲存方式完全不需要用到指標
- 建立 heap 的動作, 可以逐一 insert 的方式達成, 但是這樣耗時 O(nlogn)
- 先把所有元素不分順序直接放入 heap 中。 一開始將整個陣列的後半部看成是一大堆小 heaps, 逐層由下往上建立稍大的 heaps, 最後建立出一個完整的大 heap O(n)
- 