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
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {
	sum := 0
	countSum(root, 0, &sum)
	return sum
}

func countSum(n *TreeNode, count int, sum *int) {
	if n == nil {
		return
	}
	tmp := count*10 + n.Val
	if n.Left == nil && n.Right == nil {
		*sum += tmp
		return
	}
	countSum(n.Left, tmp, sum)
	countSum(n.Right, tmp, sum)
}

// @lc code=end
func Test(t *testing.T) {
	tc := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 25},
		{
			&TreeNode{
				4,
				&TreeNode{9, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}},
				&TreeNode{0, nil, nil},
			},
			1026,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, sumNumbers(tt.root))
	}
}
