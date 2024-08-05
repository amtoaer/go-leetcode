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
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var calculate func(*TreeNode) int
	calculate = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		left := max(calculate(n.Left), 0)
		right := max(calculate(n.Right), 0)
		res = max(res, left+right+n.Val)
		return max(left, right) + n.Val
	}
	calculate(root)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 6},
		{
			&TreeNode{
				-10,
				&TreeNode{9, nil, nil},
				&TreeNode{
					20,
					&TreeNode{15, nil, nil},
					&TreeNode{7, nil, nil},
				},
			},
			42,
		},
		{
			&TreeNode{
				1,
				&TreeNode{2, nil, nil},
				&TreeNode{3, nil, nil},
			},
			6,
		},
		{
			&TreeNode{
				-3,
				nil,
				&TreeNode{1, nil, nil},
			},
			1,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxPathSum(tt.root))
	}
}
