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

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */

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
	var res int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if k <= 0 {
			return
		}
		k--
		if k == 0 {
			res = node.Val
		}
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		k      int
		output int
	}{
		{
			input: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 4},
			},
			k:      1,
			output: 1,
		},
		{
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:      3,
			output: 3,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, kthSmallest(tt.input, tt.k), tt.output)
	}
}
