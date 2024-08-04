package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] Invert Binary Tree
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
func invertTree(root *TreeNode) *TreeNode {
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		levelNodes := nodes
		nodes = nodes[len(nodes):]
		for _, node := range levelNodes {
			if node == nil {
				continue
			}
			node.Left, node.Right = node.Right, node.Left
			nodes = append(nodes, node.Left, node.Right)
		}
	}
	return root
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output *TreeNode
	}{
		{
			&TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
			},
			&TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
			},
		},
		{nil, nil},
	}
	for _, tt := range tc {
		assert.True(t, reflect.DeepEqual(tt.output, invertTree(tt.input)))
	}
}
