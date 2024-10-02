package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
func goodNodes(root *TreeNode) int {
	return calculate(root, math.MinInt32)
}

func calculate(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	var cnt int
	if maxVal <= node.Val {
		cnt = 1
		maxVal = node.Val
	}
	return cnt + calculate(node.Left, maxVal) + calculate(node.Right, maxVal)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			output: 4,
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			output: 3,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, goodNodes(tt.input))
	}
}
