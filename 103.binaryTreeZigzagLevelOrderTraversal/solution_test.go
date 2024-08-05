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
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	var flag bool
	for len(nodes) > 0 {
		lenNodes := len(nodes)
		level := make([]int, 0, lenNodes)
		for _, node := range nodes {
			level = append(level, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		if flag {
			i, j := 0, lenNodes-1
			for i < j {
				level[i], level[j] = level[j], level[i]
				i++
				j--
			}
		}
		flag = !flag
		res = append(res, level)
		nodes = nodes[lenNodes:]
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output [][]int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			output: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, zigzagLevelOrder(tt.input))
	}
}
