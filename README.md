# golang_find_kth_element_in_BST

Given the `root` of a binary search tree, and an integer `k`, return *the* `kth` *smallest value (**1-indexed**) of all the values of the nodes in the tree*.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg](https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg)

```
Input: root = [3,1,4,null,2], k = 1
Output: 1

```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg](https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg)

```
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

```

**Constraints:**

- The number of nodes in the tree is `n`.
- `1 <= k <= n <= 104`
- `0 <= Node.val <= 104`

**Follow up:** If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

## 解析

給定一個二元搜尋樹的根結點 root 還有一個數字正整數 k

要求寫一個演算法找出排序由小至大第 k 個元素

在二元搜尋要找出第 k 小的元素

需要知道二元搜尋樹特性

給定一個二元搜尋樹根結點 root 具有以下特性:

1. 所有 root.Left 所形成的樹的結點值都小於 root.Val
2. 所有 root.Right 所形成的樹的結點值都大於 root.Val

所以只要使用 In-order Traversal 順序找到第 k 個元素就是第 k 個小的元素

In-Order Traversal 搜尋順序如下： In-Order-Traversal(root.Left), root, In-Order-Traversal(root.Right)

如下圖：

![](https://i.imgur.com/xfzWIRC.png)

會發現一定要走到最後一個 Left Leave 才開始計算順位

所以時間複雜度是 O(H+k), 其中 H 是二元搜尋樹的高度

H的值在結點很偏一邊是 N-1, 最好的狀況是 logN

題目問如說如過可以改變 BST 的結構在 insert 跟 delete 時做修改

這樣是否有辦法優化搜尋 第 k 順位的點的時間複雜度

假設除了BST 本身之外另外透過雙向連結的結構來儲存由大到小的順序

多了一個指向最大 還有最小的指標

1. 每次 insert 都找 BST 中插入的點位置,時間複雜度 O(logN)
2. 每次 delete 都找 BST 中刪除的點位置,時間複雜度 O(logN)
3. 因為有雙向鏈結每次搜尋第 k 順為 O(k)
4. 空間複雜度多存雙向鏈結是 O(N)

![](https://i.imgur.com/CgQuQnL.png)
## 程式碼

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := InOrderTraversalKth(root, k, &count)
	return result.Val
}

func InOrderTraversalKth(root *TreeNode, k int, count *int) *TreeNode {
	if root == nil { // last node
		return nil
	}
	left := InOrderTraversalKth(root.Left, k, count)
	if *count == k {
		return left
	}
	*count++
	if *count == k {
		return root
	}
	return InOrderTraversalKth(root.Right, k, count)
}

```
## 困難點

1. 理解 二元搜尋樹特性
2. 理解 in-order traversal 特性

## Solve Point

- [x]  Understand What problem need to solve
- [x]  Analysis Complexity