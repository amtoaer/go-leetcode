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
 * @lc app=leetcode.cn id=104 lang=golang
 * @lcpr version=20004
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	var (
		res int
		dfs func(*TreeNode, int)
	)
	dfs = func(tn *TreeNode, i int) {
		if tn == nil {
			return
		}
		res = max(res, i)
		dfs(tn.Left, i+1)
		dfs(tn.Right, i+1)
	}
	dfs(root, 1)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: 2,
		},
		{
			root:     nil,
			expected: 0,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		if result := maxDepth(test.root); result != test.expected {
			t.Errorf("For tree %v, expected %d but got %d", test.root, test.expected, result)
		}
	}
}
