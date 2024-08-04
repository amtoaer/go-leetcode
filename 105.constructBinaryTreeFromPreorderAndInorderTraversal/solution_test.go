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
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	loc := 0
	for inorder[loc] != rootVal {
		loc++
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:1+loc], inorder[:loc]),
		Right: buildTree(preorder[1+loc:], inorder[loc+1:]),
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		preorder []int
		inorder  []int
		output   *TreeNode
	}{
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, tt := range tc {
		assert.Equal(t, tt.output, buildTree(tt.preorder, tt.inorder))
	}
}
