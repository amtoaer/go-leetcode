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
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] Path Sum III
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

func rootSum(node *TreeNode, targetSum int) int {
	var res int
	if node == nil {
		return res
	}
	nextTarget := targetSum - node.Val
	if nextTarget == 0 {
		res++
	}
	return res + rootSum(node.Left, nextTarget) + rootSum(node.Right, nextTarget)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return rootSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		target int
		output int
	}{
		{
			input: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: -2,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: -3,
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			target: 8,
			output: 3,
		},
		{
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			target: 22,
			output: 3,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -2,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: -1,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -2,
					},
				},
			},
			target: -1,
			output: 4,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, pathSum(tt.input, tt.target))
	}
}
