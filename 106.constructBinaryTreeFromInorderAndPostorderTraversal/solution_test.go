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
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	loc := 0
	for inorder[loc] != rootVal {
		loc++
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:loc], postorder[:loc]),
		Right: buildTree(inorder[loc+1:], postorder[loc:len(postorder)-1]),
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		inorder   []int
		postorder []int
		output    *TreeNode
	}{
		{
			[]int{9, 3, 15, 20, 7},
			[]int{9, 15, 7, 20, 3},
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
	for _, test := range tc {
		assert.Equal(t, test.output, buildTree(test.inorder, test.postorder))
	}
}
