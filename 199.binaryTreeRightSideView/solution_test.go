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
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		lenNodes := len(nodes)
		for i, node := range nodes[:lenNodes] {
			if i == 0 {
				res = append(res, node.Val)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
		}
		nodes = nodes[lenNodes:]
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output []int
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, []int{1, 3}},
		{
			&TreeNode{
				1,
				&TreeNode{2, nil, &TreeNode{5, nil, nil}},
				&TreeNode{
					3,
					nil,
					&TreeNode{4, nil, nil},
				},
			},
			[]int{1, 3, 4},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, rightSideView(tt.input))
	}
}
