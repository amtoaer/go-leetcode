package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=230 lang=golang
 * @lcpr version=20004
 *
 * [230] 二叉搜索树中第 K 小的元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var items []int
	var traversal func(*TreeNode)
	traversal = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		traversal(tn.Left)
		items = append(items, tn.Val)
		traversal(tn.Right)
	}
	traversal(root)
	return items[k-1]
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,4,null,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,null,1]\n3\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		k        int
		expected int
	}{
		{&TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, 1, 1},
		{&TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}, 3, 3},
		{&TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}, 4, 4},
		{&TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}, 5, 5},
	}

	for _, test := range tests {
		if result := kthSmallest(test.root, test.k); result != test.expected {
			t.Errorf("kthSmallest(%v, %d) = %d; expected %d", test.root, test.k, result, test.expected)
		}
	}
}
