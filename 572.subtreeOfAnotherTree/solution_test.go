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
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(root *TreeNode, other *TreeNode) bool {
	if root == nil && other == nil {
		return true
	}
	if root == nil || other == nil {
		return false
	}
	return root.Val == other.Val && isSame(root.Left, other.Left) && isSame(root.Right, other.Right)
}

// @lc code=end

func Test(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	subRoot := &TreeNode{Val: 4}
	subRoot.Left = &TreeNode{Val: 1}
	subRoot.Right = &TreeNode{Val: 2}
	assert.True(t, isSubtree(root, subRoot))
}
