/*
 * @lc app=leetcode.cn id=543 lang=golang
 * @lcpr version=30201
 *
 * [543] 二叉树的直径
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var (
		res   = 1
		depth func(node *TreeNode) int
	)
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		res = max(res, left+right+1)
		return max(left, right) + 1
	}
	depth(root)
	return res - 1
}

// @lc code=end
func Test(t *testing.T) {
	tc := []struct {
		root *TreeNode
		want int
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			3,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			1,
		},
	}

	for _, tt := range tc {
		assert.Equal(t, tt.want, diameterOfBinaryTree(tt.root))
	}
}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
