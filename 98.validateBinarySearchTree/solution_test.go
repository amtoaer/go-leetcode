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
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	last := math.MinInt
	for len(stack) > 0 || root != nil {
		for node := root; node != nil; node = node.Left {
			stack = append(stack, node)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= last {
			return false
		}
		last = root.Val
		root = root.Right
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *TreeNode
		output bool
	}{
		{input: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, output: true},
		{
			input: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
			},
			output: false,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, isValidBST(tt.input), tt.output)
	}
}
