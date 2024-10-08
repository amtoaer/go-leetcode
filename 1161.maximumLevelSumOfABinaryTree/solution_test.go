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
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] Maximum Level Sum of a Binary Tree
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
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum, res := math.MinInt32, 0
	row := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		row++
		sum, length := 0, len(nodes)
		for _, node := range nodes[:length] {
			sum += node.Val
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		nodes = nodes[length:]
		if sum > maxSum {
			maxSum = sum
			res = row
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
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
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, maxLevelSum(test.root))
	}
}
