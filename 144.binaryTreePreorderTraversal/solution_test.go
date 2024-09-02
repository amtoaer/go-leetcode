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
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
	var (
		res      = []int{}
		preorder func(node *TreeNode)
	)

	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)

	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			want: []int{1, 2, 3},
		},
		{
			root: nil,
			want: []int{},
		},
		{
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, preorderTraversal(tt.root))
	}
}
