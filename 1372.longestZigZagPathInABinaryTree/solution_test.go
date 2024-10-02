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
 * @lc app=leetcode.cn id=1372 lang=golang
 *
 * [1372] Longest ZigZag Path in a Binary Tree
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

func longestZigZag(root *TreeNode) int {
	var (
		res int
		dfs func(*TreeNode, bool, int)
	)
	dfs = func(node *TreeNode, isRight bool, prev int) {
		if node == nil {
			return
		}
		res = max(res, prev+1)
		if isRight {
			dfs(node.Right, false, prev+1)
			dfs(node.Left, true, 1)
		} else {
			dfs(node.Left, true, prev+1)
			dfs(node.Right, false, 1)
		}
	}
	dfs(root, true, 0)
	dfs(root, false, 0)
	return res - 1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output int
	}{
		{nil, -1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, longestZigZag(tt.input))
	}
}
