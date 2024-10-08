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
 * @lc app=leetcode.cn id=700 lang=golang
 *
 * [700] Search in a Binary Search Tree
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
func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil && cur.Val != val {
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur
}

// @lc code=end

func Test(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	tests := []struct {
		val      int
		expected *TreeNode
	}{
		{val: 2, expected: root.Left},
		{val: 3, expected: root.Left.Right},
		{val: 5, expected: nil},
		{val: 7, expected: root.Right},
		{val: 1, expected: root.Left.Left},
	}

	for _, test := range tests {
		result := searchBST(root, test.val)
		assert.Equal(t, test.expected, result)
	}
}
