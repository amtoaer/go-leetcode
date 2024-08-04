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
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isReverse(root.Left, root.Right)
}

func isReverse(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isReverse(left.Left, right.Right) &&
		isReverse(left.Right, right.Left)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output bool
	}{
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
			},
			true,
		},
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
			false,
		},
		{nil, true},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, isSymmetric(tt.input))
	}
}
