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
 * @lc app=leetcode.cn id=2236 lang=golang
 *
 * [2236] Root Equals Sum of Children
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
func checkTree(root *TreeNode) bool {
	return root.Left.Val+root.Right.Val == root.Val
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, true, checkTree(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}))
}
