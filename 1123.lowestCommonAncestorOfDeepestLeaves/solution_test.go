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
 * @lc app=leetcode.cn id=1123 lang=golang
 *
 * [1123] Lowest Common Ancestor of Deepest Leaves
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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, node := dfs(root)
	return node
}

func dfs(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	lDepth, lNode := dfs(root.Left)
	rDepth, rNode := dfs(root.Right)
	if lDepth < rDepth {
		return rDepth + 1, rNode
	}
	if lDepth > rDepth {
		return lDepth + 1, lNode
	}
	return lDepth + 1, root
}

// @lc code=end

func Test(t *testing.T) {
	tn := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	assert.Equal(t, 2, lcaDeepestLeaves(tn).Val)
}
