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
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] Same Tree
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input1 *TreeNode
		input2 *TreeNode
		output bool
	}{
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			true,
		},
		{
			&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			false,
		},
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1},
			},
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			false,
		},
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			false,
		},
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			false,
		},
		{
			nil,
			nil,
			true,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, isSameTree(tt.input1, tt.input2))
	}
}
