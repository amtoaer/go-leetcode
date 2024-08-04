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
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output int
	}{
		{
			&TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			3,
		},
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, 2},
		{nil, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, maxDepth(tt.input))
	}
}
